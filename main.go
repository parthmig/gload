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
	results := make(chan Result, totalRequests)
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

	for result := range results {
		fmt.Printf("%+v\n", result)
	}

}
