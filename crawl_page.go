package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) (map[string]int, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return nil, err
	}
	if baseURL.Host != currentURL.Host {
		return pages, nil
	}

	currentNormalized, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return pages, err
	}

	_, ok := pages[currentNormalized]
	if ok {
		pages[currentNormalized]++
	} else {
		pages[currentNormalized] = 1
	}

	res, err := http.Get(rawCurrentURL)
	if err != nil {
		return pages, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return pages, fmt.Errorf("error fetching %s: status code %d", rawCurrentURL, res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return pages, err
	}

	htmlContent := string(data)

	fmt.Println("crawling: ", rawCurrentURL)
	result, err := getURLsFromHTML(htmlContent, rawCurrentURL)
	if err != nil {
		return pages, err
	}

	for _, page := range result {
		currentNormalized, err := normalizeURL(page)
		if err != nil {
			return pages, err
		}

		if _, ok := pages[currentNormalized]; ok {
			pages[currentNormalized]++
			continue
		}

		_, err = crawlPage(rawBaseURL, page, pages)
		if err != nil {
			fmt.Printf("error crawling %s: %v\n", page, err)
			continue
		}
	}

	return pages, nil
}
