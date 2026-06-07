# gload

`gload` is a small command-line HTTP load tester.

It sends concurrent requests to a URL and reports request counts, status codes, errors, and latency stats.

## What It Does

- Sends `N` HTTP requests to a target URL
- Limits concurrency with a configurable worker count
- Tracks successful responses and errors
- Counts HTTP status codes
- Treats non-2xx responses as errors
- Reports min, max, p50, p95, and p99 latency for successful requests

## Usage

Run with the default settings:

```sh
go run .
```

Run against a specific URL:

```sh
go run . -url https://example.com -n 100 -c 10 -timeout 10
```

## Flags

```text
-url      Target URL
-n        Total number of requests
-c        Number of concurrent requests
-timeout  Request timeout in seconds
```

Example:

```sh
go run . -url https://httpbin.org/status/200,404,500 -n 20 -c 5
```

Example output:

```text
Load Test Summary
=================

Target
------
URL:         https://httpbin.org/status/200,404,500
Requests:    20
Concurrency: 5
Timeout:     10s

Counts
------
Total:       20
Successful:  4
Errors:      16

Status Codes
------------
200:         4
404:         7
500:         9

Latency
-------
Min:         37.10ms
Max:         181.70ms
P50:         39.57ms
P95:         41.54ms
P99:         41.54ms
```

## Project Structure

```text
main.go    Program entrypoint
config.go  CLI flag parsing and validation
runner.go  Concurrent request runner
fetch.go   Single HTTP request logic
stats.go   Result aggregation and latency stats
output.go  CLI output formatting
```

## Notes

This is intentionally a small utility, not a replacement for mature load testing tools like `ab`, `hey`, or `wrk`.

I built `gload` as a hands-on way to learn Go while making something practical. It helped me practice structs, goroutines, channels, `sync.WaitGroup`, error handling, `defer`, slices, maps, sorting, CLI flags, and `net/http`.
