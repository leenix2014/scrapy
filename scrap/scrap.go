package scrap

import (
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Href struct {
	Parent string
	Name   string
}

// GetAllPdf urls=map[root][postfix], result=map[url][parent, name...]
func GetAllPdf(urls map[string]string) map[string]Href {
	allPdfs := make(map[string]Href)
	for url, postfix := range urls {
		pdfs := getPdf(url, postfix)
		for k, v := range pdfs {
			allPdfs[k] = v
		}
	}
	return allPdfs
}

func getPdf(root string, postfix string) map[string]Href {
	url, _ := url.Parse(root)
	resp, err := http.Get(root)
	if err != nil {
		log.Printf("访问%v失败%v", root, err)
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("访问%v状态错误%s", root, resp.Status)
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("%v文档解析失败%v", root, err)
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(body))
		return nil
	}

	pdfs := make(map[string]Href)
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.Contains(href, postfix) {
			var key string
			if strings.HasPrefix(href, "/") {
				key = url.Scheme + "://" + url.Host + href
			} else {
				key = url.String() + href
			}
			name := s.Text()
			pdfs[key] = Href{Parent: root, Name: name}
		}
	})
	return pdfs
}
