package main

import (
	// "fmt"

	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	htmlReader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(htmlReader)
	if err != nil {
		return ""
	}

	selected := doc.Find("h1").First().Text()
	if selected == "" {
		selected = doc.Find("h2").First().Text()
	}
	return strings.TrimSpace(selected)
}

func getFirstParagraphFromHTML(html string) string {
	htmlReader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(htmlReader)
	if err != nil {
		return ""
	}

	selected := doc.Find("main").Find("p").First().Text()
	if selected == "" {
		selected = doc.Find("p").First().Text()
	}
	return strings.TrimSpace(selected)
}
