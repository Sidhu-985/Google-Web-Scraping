package main

import (
	"fmt"
	"time"
	"github.com/gocolly/colly"
)

func main(){
    c:=colly.NewCollector()

    var searchQuery string
    fmt.Println("Enter the input:")
    fmt.Scan(&searchQuery)

    GoogleScrape(c,searchQuery)
}

func GoogleScrape(c *colly.Collector,searchQ string){
    c.SetRequestTimeout(time.Second*100)

    c.OnHTML("div.a4bIc div.YacQ span",func(h *colly.HTMLElement) {
    })

    c.OnHTML("li.sbct.PZPZIf",func(h *colly.HTMLElement) {
        result := h.ChildText(".wM6W7d span")
        fmt.Println(result)
    })

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting:",r.URL.String())
        fmt.Println("Top 10 results")
    })

    c.OnError(func(r *colly.Response, err error) {
        fmt.Println("Error: ",err.Error())
    })

    c.OnResponse(func(r *colly.Response) {
        
    })

    c.OnScraped(func(r *colly.Response) {
        fmt.Println("Finished Scraping\nStatus Code:",r.StatusCode)
    })

    c.Visit("https://www.google.com/")
}