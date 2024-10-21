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
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
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
		
		title := h.DOM.Find("h3").Text()
        link := h.Attr("href")
    
		fmt.Printf("%d.Title: %s\n",i,title)
        fmt.Println("Link:",link)
        fmt.Println()
        i++

        time.Sleep(time.Duration(100)*time.Millisecond)
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

	c.Visit("https://www.google.com/search?q="+searchQ)
}
