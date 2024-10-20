package main

import (
	"fmt"
	"time"
	"github.com/gocolly/colly"
)

func main() {
    
    c := colly.NewCollector()

    var input string
    fmt.Println("Enter the input:")
    fmt.Scan(&input)

    c.OnHTML("a[href]",func(h *colly.HTMLElement) {
        fmt.Println(h.Text)
        time.Sleep(1000*time.Millisecond)
    })

    c.OnRequest(func(r *colly.Request) {
        r.Headers.Set("Accept-Language","en-US;q=0.9")
        fmt.Println("Visiting: ",r.URL)
    })

    c.OnError(func(r *colly.Response, err error) {
        if err!=nil{
            fmt.Println("No Error Occured")
        }else{
            fmt.Println(err)
        }
    })

    c.OnScraped(func(r *colly.Response) {
        fmt.Println(r.StatusCode)
    })

    c.Visit("https://www.google.com/")
}