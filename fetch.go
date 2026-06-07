package main

import (
	"io"
	"net/http"
	"time"
)

type Result struct {
	StatusCode int
	Latency    time.Duration
	Error      error
}

func fetch(client *http.Client, url string) Result {
	start := time.Now()

	resp, err := client.Get(url)

	if err != nil {
		return Result{Error: err}
	}

	defer resp.Body.Close()

	_, err = io.Copy(io.Discard, resp.Body)
	if err != nil {
		return Result{StatusCode: resp.StatusCode, Error: err}
	}

	latency := time.Since(start)

	return Result{StatusCode: resp.StatusCode, Latency: latency, Error: nil}
}
