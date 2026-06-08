package main

import (
	"flag"
	"fmt"
	neturl "net/url"
	"strings"
	"time"
)

type Config struct {
	URL                string
	TotalRequests      int
	ConcurrentRequests int
	Timeout            time.Duration
}

func parseConfig() Config {
	url := flag.String("url", "https://jsonplaceholder.typicode.com/todos/1", "Target URL")
	totalRequests := flag.Int("n", 10, "Total number of requests")
	concurrentRequests := flag.Int("c", 3, "Number of concurrent requests")
	timeoutSeconds := flag.Int("timeout", 10, "Request timeout in seconds")

	flag.Parse()

	return Config{
		URL:                *url,
		TotalRequests:      *totalRequests,
		ConcurrentRequests: *concurrentRequests,
		Timeout:            time.Duration(*timeoutSeconds) * time.Second,
	}
}

func validateConfig(config Config) error {
	if config.TotalRequests <= 0 {
		return fmt.Errorf("n (total requests) must be greater than 0")
	}

	if config.ConcurrentRequests <= 0 {
		return fmt.Errorf("c (concurrent requests) must be greater than 0")
	}

	if config.Timeout <= 0 {
		return fmt.Errorf("timeout must be greater than 0")
	}

	if strings.TrimSpace(config.URL) == "" {
		return fmt.Errorf("url must not be empty")
	}

	parsedURL, err := neturl.Parse(config.URL)
	if err != nil || parsedURL.Host == "" {
		return fmt.Errorf("url must be a valid absolute URL")
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return fmt.Errorf("url must use http or https")
	}

	if config.ConcurrentRequests > config.TotalRequests {
		return fmt.Errorf("c (concurrent requests) cannot be greater than n (total requests)")
	}

	return nil
}
