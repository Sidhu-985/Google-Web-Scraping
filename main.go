package main

import (
	"fmt"
	"net/url"
	"time"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.google.com"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.125 Safari/537.36"),
	)

	var searchQuery string
	fmt.Println("Enter the search query:")
	fmt.Scanln(&searchQuery)

	searchQuery = url.QueryEscape(searchQuery)

	GoogleScrape(c, searchQuery)
}

func GoogleScrape(c *colly.Collector, searchQ string) {
	c.SetRequestTimeout(100 * time.Second)

	i:=1
	c.OnHTML("div.yuRUbf a", func(h *colly.HTMLElement) {
		if i>10 {
			return
		}

		result := h.DOM.Find("h3").Text()
		link := h.Attr("href")

		fmt.Println("===================================")
		fmt.Printf("%d. Result: %s\n",i,result)
		fmt.Println("Link:", link)
		fmt.Println("===================================")
		fmt.Println()
		i++

		time.Sleep(time.Duration(500) * time.Millisecond)
	})

	c.OnHTML("a#pnnext", func(h *colly.HTMLElement) {
		if i<11 {
			nextPage := h.Attr("href")
			c.Visit("https://www.google.com"+nextPage)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
		fmt.Println("Top 10 results:")
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error:", err.Error())
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished Scraping.\nStatus Code:", r.StatusCode)
	})

	c.Visit("https://www.google.com/search?q=" + searchQ)
}
