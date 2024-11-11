package main

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, err
	}

	hrefs := []string{}
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attribute := range n.Attr {
				if attribute.Key == "href" {
					hrefs = append(hrefs, attribute.Val)
					break
				}
			}
		}
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			traverse(child)
		}
	}
	traverse(doc)

	urls := []string{}
	for _, nl := range urls {
		parsed, err := url.Parse(nl)
		if err != nil {
			return []string{}, err
		}

		if parsed.Host == "" {
			urls = append(urls, rawBaseURL+parsed.Path)
		} else {
			urls = append(urls, parsed.Scheme+parsed.Host+parsed.Path)
		}
	}

	return urls, nil
}
