package search

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

const defaultAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"

const baseGoogleUrl = "https://www.google.com/search?"

func Search(params SearchParameters) *SearchResult {
	c := colly.NewCollector()
	c.UserAgent = defaultAgent

	c.OnHTML("div.g", func(e *colly.HTMLElement) {
		sel := e.DOM

		linkHref, _ := sel.Find("a").Attr("href")
		linkText := strings.TrimSpace(linkHref)
		titleText := strings.TrimSpace(sel.Find("div > div > div > div > span > a > h3").Text())
		descText := strings.TrimSpace(sel.Find("div > div > div > div:first-child > span:first-child").Text())
		fmt.Println("=============OnHTML===============")
		fmt.Println(linkText)
		fmt.Println(titleText)
		fmt.Println(descText)
		fmt.Println("=============OnHTML===============")
		fmt.Println()
	})

	url := buildUrl(params)
	c.Visit(url)

	return nil
}

func buildUrl(params SearchParameters) string {
	url := baseGoogleUrl + "q=" + params.SearchString
	return url
}
