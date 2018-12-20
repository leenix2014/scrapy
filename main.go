package main

import (
	"scrapy/mail"
	"scrapy/scrap"
	"strings"
)

func main() {
	urls := []string{
		"http://www.cmbc.com.cn/jrms/msdt/yjbg/index.htm",
		"https://www.hkma.gov.hk/gb_chi/publications-and-research/annual-report/2017.shtml",
		"https://www.bochk.com/m/sc/investment/econanalysis/bocecon.html",
	}

	allNew := make(map[string]bool)
	for _, url := range urls {
		newPdfs := scrap.GetNewPdf(url)
		for k, v := range newPdfs {
			allNew[k] = v
		}
	}

	cu := strings.Split(scrap.CurrentUser, "@")
	body := "Dear " + cu[0] + ", \n\n 检测到有以下更新pdf: \n"
	for pdf, _ := range allNew {
		body += "<a href=\"" + pdf + "\">" + pdf + "<a/>\n<br>"
	}
	mail.SendHtml(scrap.CurrentUser+";test@liquanlin.tech", "文档有更新", body)
}
