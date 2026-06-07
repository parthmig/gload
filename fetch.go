package main

import (
	"io"
	"net/http"
	"time"
)

type Result struct {
	StatusCode int
	LatencyMs  time.Duration
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
		return Result{Error: err}
	}
	// close the response body
	defer resp.Body.Close()
	// discard the response body
	_, err = io.Copy(io.Discard, resp.Body)
	if err != nil {
		return Result{StatusCode: resp.StatusCode, Error: err}
	}
	// calculate latency
	latency := time.Since(start)

	return Result{StatusCode: resp.StatusCode, LatencyMs: latency, Error: nil}
}
