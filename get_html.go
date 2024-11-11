package main

import (
	"fmt"
	"io"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}

	if status := res.StatusCode; status > 399 {
		return "", fmt.Errorf("http error code: %d", status)
	}

	if contentType := res.Header.Get("Content-Type"); contentType != "text/html" {
		return "", fmt.Errorf("website returned non-html content: %v", contentType)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
