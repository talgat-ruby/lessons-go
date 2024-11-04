package main

import (
	"fmt"
	"time"
)

// Exercise: Concurrent URL Fetcher with WaitGroup
//
// Goal: Implement a program that fetches data from multiple URLs concurrently,
// then waits for all fetch operations to complete before proceeding.
//
// · Simulate fetching data from a list of URLs using goroutines.
// · Add save the result to results slice

func fetcher(id int, url string) string {
	time.Sleep(time.Duration(id+1) * 1 * time.Second)
	return fmt.Sprintf("Response from %s", url)
}

func main() {
	urls := []string{
		"http://example.com/page1",
		"http://example.com/page2",
		"http://example.com/page3",
		"http://example.com/page4",
	}
	results := make([]string, 0, len(urls))

	fmt.Println("All fetches completed. Results:")
	for _, result := range results {
		fmt.Println(result)
	}
}
