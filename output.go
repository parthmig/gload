package main

import (
	"fmt"
	"time"
)

func formatDuration(d time.Duration) string {
	return fmt.Sprintf("%.2fms", float64(d.Microseconds())/1000)
}

func printSummary(summary Stats) {
	fmt.Println("Load Test Summary")
	fmt.Println("-------")
	fmt.Printf("Total requests: %d\n", summary.Total)
	fmt.Printf("Successful:     %d\n", summary.Successes)
	fmt.Printf("Errors:         %d\n", summary.Errors)
	fmt.Printf("Min latency:    %s\n", formatDuration(summary.Min))
	fmt.Printf("Max latency:    %s\n", formatDuration(summary.Max))
	fmt.Printf("P50 latency:    %s\n", formatDuration(summary.P50))
	fmt.Printf("P95 latency:    %s\n", formatDuration(summary.P95))
	fmt.Printf("P99 latency:    %s\n", formatDuration(summary.P99))
}
