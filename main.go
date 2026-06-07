package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {

	config := parseConfig()

	if err := validateConfig(config); err != nil {
		fmt.Println("Error:", err)
		return
	}

	client := &http.Client{
		Timeout: config.Timeout,
	}

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

	summary := calculateStats(allResults)
	printSummary(summary)

}
