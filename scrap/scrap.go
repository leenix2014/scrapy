package scrap

import (
	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"scrapy/entity"
	"strings"
)

var users = make(map[string]map[string]bool) //map[user_mail][pdf]visited
var engine *xorm.Engine
var CurrentUser = "liquanlin@liquanlin.tech"

func init() {
	loadFromDB()
}

func loadFromDB() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:root@(127.0.0.1:3306)/lql?charset=utf8")
	if err != nil {
		log.Fatalf("无法连接数据库%s", err)
	}
	var dbs []entity.TPdf
	engine.Table(entity.TPdf{}).Find(&dbs)

	for _, db := range dbs {
		pdfs, exist := users[db.UserMail]
		if !exist {
			pdfs = make(map[string]bool)
			users[db.UserMail] = pdfs
		} else {
			pdfs[db.Url] = db.Visited
		}
	}
}

func GetNewPdf(root string) map[string]bool {
	url, _ := url.Parse(root)
	resp, err := http.Get(root)
	if err != nil {
		log.Printf("访问%v失败", root)
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("访问%v状态错误%s", root, resp.Status)
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("%v文档解析失败", root)
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(body))
		return nil
	}

	pdfs, exists := users[CurrentUser]
	if !exists {
		pdfs = make(map[string]bool)
		users[CurrentUser] = pdfs
	}
	newPdfs := make(map[string]bool)
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.Contains(href, ".pdf") {
			var key string
			if strings.HasPrefix(href, "/") {
				key = url.Scheme + "://" + url.Host + href
			} else {
				key = url.String() + href
			}
			_, exists := pdfs[key]
			if !exists {
				newPdfs[key] = false
				pdfs[key] = false
			}
		}
	})
	for key, visited := range newPdfs {
		bean := entity.TPdf{UserMail: CurrentUser, Root: root, Url: key, Visited: visited}
		_, err := engine.InsertOne(bean)
		if err != nil {
			log.Printf("%v插入失败(%v)", bean, err)
		}
	}
	return newPdfs
}
