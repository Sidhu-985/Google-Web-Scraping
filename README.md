# Google Web Scraper in Go

This project is a simple web scraper built in Go using the `Colly` library to scrape the top 10 Google search results for a given user input. It extracts the titles and links of the search results.

**Note**: Scraping Google directly may violate Google's Terms of Service and can lead to IP blocking. For production purposes, refer to using this [Google Custom Search JSON API](https://developers.google.com/custom-search/v1/introduction) instead.

## Features
- The scraper sends a request to Google with the given keyword and capture the
search results page. It will then parse the HTML content and extract the top 10 search
result in a readable manner.

## Prerequisites

- Go 1.16 or later installed on your machine.
- The `Colly` library for web scraping. Install it by running:
  ```go
  go get -u github.com/gocolly/colly

## To run the code:
```go
go run main.go
