package main

import (
	"errors"
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
		return "", fmt.Errorf("error code: %d", status)
	}

	if contentType := res.Header.Get("content-type"); contentType != "text/html" {
		return "", errors.New("website returned non-html content")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
