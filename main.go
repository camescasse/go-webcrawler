package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	html, err := getHTML(args[0])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(html)
	fmt.Printf("starting crawl of: %s\n", args[0])
	os.Exit(0)
}
