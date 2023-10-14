package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_ScrapeURL(t *testing.T) {
	// Create a test server to serve sample HTML content
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		htmlContent := `
            <html>
                <head>
                    <title>Test Page</title>
                </head>
                <body>
                    <a href="https://example.com">Example</a>
                    <a href="https://example.org">Another Example</a>
                </body>
            </html>
        `
		w.Write([]byte(htmlContent))
	}))
	defer server.Close()

	url := server.URL

	results := make(chan PageData)
	go scrapeURL(url, results)
	data := <-results

	expected := PageData{
		Title: "Test Page",
		Links: []string{"https://example.com", "https://example.org"},
	}

	if data.Title != expected.Title || data.Links[0] != expected.Links[0] || data.Links[1] != expected.Links[1] {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}

func Test_ScrapeURL_InvalidURL(t *testing.T) {
	url := "https://invalidurlthatdoesnotexist.com"

	results := make(chan PageData)
	go scrapeURL(url, results)
	data := <-results

	if data.Title != "" || len(data.Links) != 0 {
		t.Errorf("Expected empty PageData, but got %v", data)
	}
}

func Test_ScrapeURL_InvalidStatus(t *testing.T) {
	// Create a test server that returns a non-OK status
	server := httptest.NewServer(http.NotFoundHandler())
	defer server.Close()

	url := server.URL

	results := make(chan PageData)
	go scrapeURL(url, results)
	data := <-results

	if data.Title != "" || len(data.Links) != 0 {
		t.Errorf("Expected empty PageData, but got %v", data)
	}
}
