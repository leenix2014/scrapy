package scrap

import (
	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func GetAllPdf(root string) map[string]string {
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

	pdfs := make(map[string]string)
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.Contains(href, ".pdf") {
			var key string
			if strings.HasPrefix(href, "/") {
				key = url.Scheme + "://" + url.Host + href
			} else {
				key = url.String() + href
			}
			pdfs[key] = root
		}
	})
	return pdfs
}
