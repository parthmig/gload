package main

import (
	"slices"
	"time"
)

type Stats struct {
	Total     int
	Successes int
	Errors    int
	Min       time.Duration
	Max       time.Duration
	P50       time.Duration
	P95       time.Duration
	P99       time.Duration
}

func calculateStats(results []Result) Stats {
	totalRequests := len(results)
	totalErrors := 0
	totalSuccesses := 0
	var latencies []time.Duration

	for _, result := range results {
		if result.Error != nil {
			totalErrors++
			continue
		}

		totalSuccesses++
		latencies = append(latencies, result.Latency)
	}

	slices.Sort(latencies)

	// if all requests failed, return counts only
	if len(latencies) == 0 {
		return Stats{
			Total:  totalRequests,
			Errors: totalErrors,
		}
	}

	min := latencies[0]
	max := latencies[len(latencies)-1]
	p50 := latencies[int(float64(len(latencies)-1)*0.50)]
	p95 := latencies[int(float64(len(latencies)-1)*0.95)]
	p99 := latencies[int(float64(len(latencies)-1)*0.99)]

	return Stats{
		Total:     totalRequests,
		Successes: totalSuccesses,
		Errors:    totalErrors,
		Min:       min,
		Max:       max,
		P50:       p50,
		P95:       p95,
		P99:       p99,
	}
}
