package logic

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"scrapy/entity"
	"scrapy/mail"
	"scrapy/scrap"
	"scrapy/util"
	"strings"
)

var users = make(map[string]map[string]bool) //map[user_mail][pdf]visited
var engine *xorm.Engine
var currentUser = "349382785@qq.com"

func init() {
	loadFromDB()
}

func loadFromDB() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:root@(127.0.0.1:3306)/lql?charset=utf8")
	if err != nil {
		log.Fatalf("无法连接数据库%s", err)
	}
	//engine.ShowSQL(true)
	var dbs []entity.TPdf
	engine.Table(entity.TPdf{}).Find(&dbs)
	for _, db := range dbs {
		pdfs, exist := users[db.UserMail]
		if !exist {
			pdfs = make(map[string]bool)
			users[db.UserMail] = pdfs
		}
		pdfs[db.Url] = util.ToBool(db.Visited)
	}
}

func Check() {
	log.Println("开始检查url更新")
	urls := []string{
		"http://www.cmbc.com.cn/jrms/msdt/yjbg/index.htm",
		"https://www.hkma.gov.hk/gb_chi/publications-and-research/annual-report/2017.shtml",
		"https://www.bochk.com/m/sc/investment/econanalysis/bocecon.html",
	}

	allPdfs := make(map[string]string)
	for _, url := range urls {
		pdfs := scrap.GetAllPdf(url)
		for k, v := range pdfs {
			allPdfs[k] = v
		}
	}

	user := users[currentUser]
	if user == nil {
		user = make(map[string]bool)
		users[currentUser] = user
	}
	nonVisited := make(map[string]string)
	for k, root := range allPdfs {
		visited, exists := user[k]
		if !exists {
			//更新数据库
			bean := entity.TPdf{UserMail: currentUser, Root: root, Url: k, Visited: 0}
			_, err := engine.InsertOne(bean)
			if err != nil {
				log.Printf("插入失败(%v)，Bean(%v)", err, bean)
			}
			user[k] = false //更新内存
		}
		if !visited {
			nonVisited[k] = root
		}
	}
	if len(nonVisited) == 0 {
		log.Println("无更新")
		return
	}

	cu := strings.Split(currentUser, "@")
	body := "Dear " + cu[0] + ", \n\n 检测到有以下更新pdf: \n"
	for pdf, _ := range nonVisited {
		body += "<a href=\"" + pdf + "\">" + pdf + "<a/>\n<br/><br/>"
	}
	err := mail.SendHtml(currentUser+";test@liquanlin.tech", "文档有更新", body)
	if err != nil {
		log.Printf("发送邮件失败(%v)!", err)
		return
	}
	log.Printf("已发送邮件:\n %s", body)
	for key, _ := range nonVisited {
		bean := entity.TPdf{UserMail: currentUser, Url: key}
		_, err = engine.Update(entity.TPdf{Visited: 1}, bean)
		if err != nil {
			log.Printf("更新失败(%v)，Bean(%v)", err, bean)
		} else {
			user[key] = true
		}
	}
}
