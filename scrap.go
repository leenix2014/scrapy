package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-xorm/xorm"
	"log"
	"net/http"
	"net/url"
	"scrapy/entity"
	"strings"
)

func init() {
	engine, err := xorm.NewEngine("mysql", "root:root@(127.0.0.1:3306)/lql?charset=utf8")
	if err != nil {
		log.Fatalf("无法连接数据库%s", err)
	}
	engine.Get(entity.TPdf{})
}

func main() {
	urls := []string{
		"http://www.cmbc.com.cn/jrms/msdt/yjbg/index.htm",
		"https://www.hkma.gov.hk/gb_chi/publications-and-research/annual-report/2017.shtml",
		"https://www.bochk.com/m/sc/investment/econanalysis/bocecon.html",
	}

	for _, url := range urls {
		getAllPdf(url)
	}
}

func getAllPdf(root string) {
	url, _ := url.Parse(root)
	// Request the HTML page.
	resp, err := http.Get(root)
	if err != nil {
		//log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		//log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
		return
	}

	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	pdfs := make(map[string]bool)
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.Contains(href, ".pdf") {
			if strings.HasPrefix(href, "/") {
				pdfs[url.Scheme+"://"+url.Host+href] = false
			} else {
				pdfs[url.String()+href] = false
			}
		}
	})
	for pdf, _ := range pdfs {
		fmt.Println(pdf)
	}
}
