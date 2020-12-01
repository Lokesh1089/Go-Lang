package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	parse, err := url.Parse("https://www.instagram.com/person?name=ajay_lokesh&phone=%2B9655001234&phone=%2B123456789")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Scheme: ", parse.Scheme)
	fmt.Println("Host: ", parse.Host)

	queries := parse.Query()
	fmt.Println("Query Strings: ")
	for k, v := range queries {
		fmt.Printf("  %v = %v\n", k, v)
	}
	fmt.Println("Path: ", parse.Path)
}