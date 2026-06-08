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

Build a local binary:

```sh
go build -o gload
```

Run the binary:

```sh
./gload -url https://example.com -n 100 -c 10 -timeout 10
```

## Flags

```text
-url      Target URL
-n        Total number of requests
-c        Number of concurrent requests
-timeout  Request timeout in seconds
```

The URL must be an absolute `http` or `https` URL. Concurrency must be greater than `0` and cannot be greater than the total number of requests.

Example:

```sh
go run . -url https://httpbin.org/status/200,404,500 -n 20 -c 5
```

Example output:

Because the endpoint returns one of several status codes, exact counts and latency values may vary.

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
Successful:  6
Errors:      14

Status Codes
------------
200:         6
404:         9
500:         5

Latency
-------
Min:         34.28ms
Max:         214.53ms
P50:         35.71ms
P95:         214.53ms
P99:         214.53ms
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

## Testing

Run the test suite:

```sh
go test ./...
```

Format the code:

```sh
gofmt -w .
```

## License

MIT

## Notes

This is intentionally a small utility, not a replacement for mature load testing tools like `ab`, `hey`, or `wrk`.

I built `gload` as a hands-on way to learn Go while making something practical. It helped me practice structs, goroutines, channels, `sync.WaitGroup`, error handling, `defer`, slices, maps, sorting, CLI flags, and `net/http`.
