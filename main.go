package main

import (
	"flag"
	"net/http"
	"sync"
	"time"
)

func main() {

	url := flag.String("url", "https://jsonplaceholder.typicode.com/todos/1", "Target URL")
	totalRequests := flag.Int("n", 10, "Total number of requests")
	concurrentRequests := flag.Int("c", 3, "Number of concurrent requests")
	timeoutSeconds := flag.Int("timeout", 10, "Request timeout in seconds")

	flag.Parse()

	client := &http.Client{
		Timeout: time.Duration(*timeoutSeconds) * time.Second,
	}

	var wg sync.WaitGroup

	// results channel stores the results of requests &
	// sem limits the number of concurrent requests
	results := make(chan Result, *concurrentRequests)
	sem := make(chan struct{}, *concurrentRequests)

	for i := 0; i < *totalRequests; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			sem <- struct{}{}
			defer func() { <-sem }()

			result := fetch(client, *url)
			results <- result
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()
	var allResults []Result
	for result := range results {
		allResults = append(allResults, result)
	}

	summary := calculateStats(allResults)
	printSummary(summary)

}
