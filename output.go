package main

import (
	"fmt"
	"slices"
	"time"
)

func formatDuration(d time.Duration) string {
	return fmt.Sprintf("%.2fms", float64(d.Microseconds())/1000)
}

func printSummary(config Config, summary Stats) {
	fmt.Println("Load Test Summary")
	fmt.Println("=================")

	fmt.Println()
	fmt.Println("Target")
	fmt.Println("------")
	fmt.Printf("URL:         %s\n", config.URL)
	fmt.Printf("Requests:    %d\n", config.TotalRequests)
	fmt.Printf("Concurrency: %d\n", config.ConcurrentRequests)
	fmt.Printf("Timeout:     %s\n", config.Timeout)

	fmt.Println()
	fmt.Println("Counts")
	fmt.Println("------")
	fmt.Printf("Total:       %d\n", summary.Total)
	fmt.Printf("Successful:  %d\n", summary.Successes)
	fmt.Printf("Errors:      %d\n", summary.Errors)

	fmt.Println()
	fmt.Println("Status Codes")
	fmt.Println("------------")
	if len(summary.StatusCounts) == 0 {
		fmt.Println("No HTTP status codes recorded. Requests likely failed before receiving a response.")
	} else {
		var codes []int
		for code := range summary.StatusCounts {
			codes = append(codes, code)
		}

		slices.Sort(codes)

		for _, code := range codes {
			fmt.Printf("%d:         %d\n", code, summary.StatusCounts[code])
		}
	}

	fmt.Println()
	fmt.Println("Latency")
	fmt.Println("-------")
	fmt.Printf("Min:         %s\n", formatDuration(summary.Min))
	fmt.Printf("Max:         %s\n", formatDuration(summary.Max))
	fmt.Printf("P50:         %s\n", formatDuration(summary.P50))
	fmt.Printf("P95:         %s\n", formatDuration(summary.P95))
	fmt.Printf("P99:         %s\n", formatDuration(summary.P99))
}
