package search

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

const defaultAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"

const baseGoogleUrl = "https://www.google.com/search?"

func Search(params SearchParameters) ([]SearchResult, error) {
	c := colly.NewCollector()
	c.UserAgent = defaultAgent

	results := []SearchResult{}
	var rErr error
	var index int
	c.OnHTML("div.g", func(e *colly.HTMLElement) {
		sel := e.DOM

		linkHref, _ := sel.Find("a").Attr("href")
		linkText := strings.TrimSpace(linkHref)
		titleText := strings.TrimSpace(sel.Find("div > div > div > div > span > a > h3").Text())
		descText := strings.TrimSpace(sel.Find("div > div > div > div:first-child > span:first-child").Text())
		index++
		if linkText != "" && linkText != "#" && titleText != "" {
			result := SearchResult{
				Index:       index,
				URL:         linkText,
				Title:       titleText,
				Description: descText,
			}
			results = append(results, result)
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		rErr = err
	})

	url := buildUrl(params)
	c.Visit(url)

	if rErr != nil {
		return nil, rErr
	}

	return results, nil
}

func buildUrl(params SearchParameters) string {
	url := baseGoogleUrl + "q=" + params.SearchString
	return url
}
