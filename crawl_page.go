package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (config *config) crawlPage(currentURL *url.URL) {
	if config.baseURL.Host != currentURL.Host {
		return
	}

	currentNormalized, err := normalizeURL(currentURL.String())
	if err != nil {
		fmt.Println("error normalizing current url:", err)
		return
	}

	if _, ok := config.pages[currentNormalized]; ok {
		config.pages[currentNormalized]++
		return
	} else {
		config.pages[currentNormalized] = 1
	}

	res, err := http.Get(currentURL.String())
	if err != nil {
		fmt.Println("error getting http body:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		fmt.Printf("error fetching %s: status code: %d\n", currentNormalized, res.StatusCode)
		return
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading data:", err)
		return
	}

	htmlContent := string(data)

	fmt.Println("crawling:", currentNormalized)
	foundURLs, err := getURLsFromHTML(htmlContent, currentURL)
	if err != nil {
		fmt.Println("error crawling site:", err)
		return
	}

	for _, page := range foundURLs {
		parsedPage, err := url.Parse(page)
		if err != nil {
			fmt.Println("error parsing url:", err)
			return
		}
		config.crawlPage(parsedPage)
	}
}
