package logic

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
	"log"
	"scrapy/entity"
	"scrapy/mail"
	"scrapy/scrap"
	"scrapy/util"
	"strings"
)

var users = make(map[string]map[string]bool) //map[user_mail][pdf]visited
var engine *xorm.Engine

func Init() {
	loadFromDB()
}

func loadFromDB() {
	var err error
	dbUrl := viper.GetString("dbUrl")
	engine, err = xorm.NewEngine("mysql", dbUrl)
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
	urls := viper.GetStringSlice("watchUrls")
	allPdfs := scrap.GetAllPdf(urls)
	currentUsers := viper.GetStringSlice("watcherMails")
	for _, currentUser := range currentUsers {
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
			log.Printf("%v无更新", currentUser)
			continue
		}

		cu := strings.Split(currentUser, "@")
		body := fmt.Sprintf("Dear %v, <br/> &emsp;检测到有以下更新pdf: <br/><br/>", cu[0])
		for pdf, _ := range nonVisited {
			parts := strings.Split(pdf, "/")
			name := parts[len(parts)-1]
			body += fmt.Sprintf("<a href=\"%v\">%v<a/><br/><br/>", pdf, name)
		}
		err := mail.SendHtml(currentUser+";test@liquanlin.tech", "文档有更新", body)
		if err != nil {
			log.Printf("%v发送邮件失败(%v)!", currentUser, err)
			continue
		}
		log.Printf("已发送邮件:\n %s", body)
		for key, _ := range nonVisited {
			bean := entity.TPdf{UserMail: currentUser, Url: key}
			_, err = engine.Update(entity.TPdf{Visited: 1}, bean)
			if err != nil {
				log.Printf(""+
					"更新失败(%v)，Bean(%v)", err, bean)
			} else {
				user[key] = true
			}
		}
	}
}
