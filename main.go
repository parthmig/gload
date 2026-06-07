package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	totalRequests := 10
	concurrentRequests := 3
	url := "https://www.google.com"

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	var wg sync.WaitGroup

	// results channel stores the results of requests &
	// sem limits the number of concurrent requests
	results := make(chan Result, concurrentRequests)
	sem := make(chan struct{}, concurrentRequests)

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			sem <- struct{}{}
			defer func() { <-sem }()

			result := fetch(client, url)
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

	fmt.Printf("collected %d results\n", len(allResults))

}
