package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type PageData struct {
	Title string
	Links []string
}

func main() {
	urls := []string{
		"https://books.toscrape.com/",
		"https://www.yelp.com/",
		"https://www.reddit.com/",
		"https://en.wikipedia.org/wiki/Main_Page",
		"https://finance.yahoo.com/",
		"https://quotes.toscrape.com/",
		"https://crawler-test.com/",
		"https://realpython.github.io/fake-jobs/",
		"https://www.amazon.in/",
		"https://www.walmart.com/",
		"https://www.ebay.com/",
		"https://www.goal.com/en-in",
		"https://www.foxnews.com/",
		"https://www.bbc.com/",
		"https://www.aljazeera.com/",
		"https://abcnews.go.com/",
		"https://www.livescore.com/en/",
		"https://www.livescore.in/",
		"https://www.cricbuzz.com/",
		"https://www.flashscore.com/",
		"https://www.flipkart.com/",
		"https://soundcloud.com/discover",
	}

	results := make(chan PageData, len(urls))

	for _, url := range urls {
		go scrapeURL(url, results)
	}

	for i := 0; i < len(urls); i++ {
		data := <-results
		fmt.Println()
		fmt.Printf("Title: %s\n", data.Title)
		fmt.Printf("Links: %v\n", data.Links)
		fmt.Println("\n---------------------------------------------------------------------------------------------------------------")
		fmt.Println()
	}
}

func scrapeURL(url string, results chan PageData) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error scraping %s: %v\n", url, err)
		results <- PageData{}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error scraping %s: Status Code %d\n", url, resp.StatusCode)
		results <- PageData{}
		return
	}

	tokenizer := html.NewTokenizer(resp.Body)
	pageData := PageData{}

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			results <- pageData
			return
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.Data == "title" {
				tokenizer.Next()
				pageData.Title = strings.TrimSpace(tokenizer.Token().Data)
			} else if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						if strings.HasPrefix(attr.Val, "http") {
							pageData.Links = append(pageData.Links, attr.Val)
						}
					}
				}
			}
		}
	}
}
