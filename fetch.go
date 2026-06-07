package main

import (
	"net/http"
	"time"
)

type Result struct {
	StatusCode int
	LatencyMs  float64
	Error      error
}

func fetch(url string) Result {
	// create a new http client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	// start timer
	start := time.Now()
	// send the request
	resp, err := client.Get(url)
	// if there is an error, return the error and latency as 0
	if err != nil {
		return Result{Error: err, LatencyMs: 0}
	}
	defer resp.Body.Close()
	// calculate latency
	latencyMs := time.Since(start).Seconds() * 1000
	return Result{StatusCode: resp.StatusCode, LatencyMs: latencyMs, Error: nil}
}
