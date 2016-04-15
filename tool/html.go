package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"util"

	"github.com/PuerkitoBio/goquery"
	"github.com/lhboy1984/godatayes/helper"
	"golang.org/x/net/html"
)

func pagelinks(page string) {
	url := fmt.Sprintf("https://api.wmcloud.com/docs/plugins/pagetree/naturalchildren.action?decorator=none&excerpt=false&sort=position&reverse=false&disableLinks=false&expandCurrent=true&hasRoot=true&pageId=%s&treeId=0&startDepth=0&mobile=false&treePageId=%s&_=%d", page, page, time.Now().Unix())
	doc, err := goquery.NewDocument(url)
	helper.IfErrPanic(err)

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		if id, ok := s.Attr("id"); ok {
			reg := regexp.MustCompile(".*(\\d+)-0")
			page := reg.FindAllStringSubmatch(id, -1)[0][1]
			pagelinks(page)
		} else if href, ok := s.Attr("href"); ok {
			apilink(href, s.Text())
		}
	})
}

func links() {
	doc, err := goquery.NewDocument("https://api.wmcloud.com/docs/plugins/pagetree/naturalchildren.action?decorator=none&excerpt=false&sort=position&reverse=false&disableLinks=false&expandCurrent=true&hasRoot=true&pageId=1867793&treeId=0&startDepth=0&mobile=false&ancestors=2392676&ancestors=1867793&treePageId=1867781")
	helper.IfErrPanic(err)

	doc.Find("div#children1867781-0").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(ii int, ss *goquery.Selection) {
			if id, ok := ss.Attr("id"); ok {
				reg := regexp.MustCompile(".*(\\d+)-0")
				page := reg.FindAllStringSubmatch(id, -1)[0][1]
				pagelinks(page)
			} else if href, ok := ss.Attr("href"); ok {
				apilink(href, ss.Text())
			}
		})
	})
}

func getHtmlText(node *html.Node) string {
	if node.Type == html.TextNode {
		return node.Data
	} else if node.Type == html.ElementNode {
		for nc := node.FirstChild; nc != nil; nc = nc.NextSibling {
			if nc.Type == html.ElementNode {
				return getHtmlText(nc)
			}
		}

		for nc := node.FirstChild; nc != nil; nc = nc.NextSibling {
			if nc.Type == html.TextNode {
				return getHtmlText(nc)
			}
		}
	}

	return ""
}

func apiArgs(s *goquery.Selection) []apiField {
	var args []apiField
	nodes := s.Find("td").Nodes

	for i := 0; i < len(nodes); i = i + 4 {
		var field apiField
		field.name = getHtmlText(nodes[i])
		field.vtype = getHtmlText(nodes[i+1])
		field.empty = !(getHtmlText(nodes[i+2]) == "否")
		field.comment = getHtmlText(nodes[i+3])
		args = append(args, field)
	}

	return args
}

func apiRets(s *goquery.Selection) []apiField {
	var rets []apiField
	nodes := s.Find("td").Nodes

	for i := 0; i < len(nodes); i = i + 3 {
		var field apiField
		field.name = getHtmlText(nodes[i])
		field.vtype = getHtmlText(nodes[i+1])
		field.comment = getHtmlText(nodes[i+2])
		rets = append(rets, field)
	}

	return rets
}

func apiIndices(s *goquery.Selection) map[string]string {
	indices := map[string]string{}
	nodes := s.Find("td").Nodes

	for i := 0; i < len(nodes); i = i + 2 {
		//indices[getHtmlText(nodes[i])] = getHtmlText(nodes[i+1])
	}

	return indices
}

func apilink(url string, meta string) {
	doc, err := goquery.NewDocument("https://api.wmcloud.com" + url)
	util.IfErrPanic(err)

	apis[meta] = map[string]apiMeta{}

	doc.Find("div#main-content > .table-wrap > table.confluenceTable > tbody").Each(func(i int, s *goquery.Selection) {
		if strings.Index(s.Text(), "GET") < 0 {
			return
		}
		var api apiMeta
		sel := s.ChildrenFiltered("tr")
		sel.Each(func(ii int, ss *goquery.Selection) {
			switch ii {
			case 0:
				reg := regexp.MustCompile("/api/[^/]+/([^/]+).json")
				apiurl := reg.FindAllStringSubmatch(ss.Text(), -1)
				api.url = apiurl[0][0]
				api.name = apiurl[0][1]
			case 1:
				api.desc = ss.Text()
			case 2:
			case 3:
				ss.Find("div.table-wrap > table.confluenceTable > tbody").Each(func(iii int, sss *goquery.Selection) {
					switch iii {
					case 0:
						api.args = apiArgs(sss)
					case 1:
						api.rets = apiRets(sss)
					case 2:
						api.indices = apiIndices(sss)
					}
				})
			}
		})

		apis[meta][api.name] = api
	})
}
