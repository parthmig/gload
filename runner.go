package main

import (
	"net/http"
	"sync"
)

func runLoadTest(client *http.Client, config Config) []Result {
	var wg sync.WaitGroup

	// results channel stores the results of requests &
	// sem limits the number of concurrent requests
	results := make(chan Result, config.ConcurrentRequests)
	sem := make(chan struct{}, config.ConcurrentRequests)

	for i := 0; i < config.TotalRequests; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			sem <- struct{}{}
			defer func() { <-sem }()

			result := fetch(client, config.URL)
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

	return allResults
}
