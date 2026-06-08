package main

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestCalculateStats(t *testing.T) {
	tests := []struct {
		name string
		in   []Result
		want Stats
	}{
		{
			name: "successful responses",
			in: []Result{
				{StatusCode: 200, Latency: 10 * time.Millisecond},
				{StatusCode: 200, Latency: 20 * time.Millisecond},
				{StatusCode: 200, Latency: 30 * time.Millisecond},
				{StatusCode: 200, Latency: 40 * time.Millisecond},
			},
			want: Stats{
				Total:        4,
				Successes:    4,
				Errors:       0,
				StatusCounts: map[int]int{200: 4},
				Min:          10 * time.Millisecond,
				Max:          40 * time.Millisecond,
				P50:          20 * time.Millisecond,
				P95:          40 * time.Millisecond,
				P99:          40 * time.Millisecond,
			},
		},
		{
			name: "non 2xx responses count as errors",
			in: []Result{
				{StatusCode: 200, Latency: 10 * time.Millisecond},
				{StatusCode: 404, Latency: 20 * time.Millisecond},
				{StatusCode: 500, Latency: 30 * time.Millisecond},
			},
			want: Stats{
				Total:        3,
				Successes:    1,
				Errors:       2,
				StatusCounts: map[int]int{200: 1, 404: 1, 500: 1},
				Min:          10 * time.Millisecond,
				Max:          10 * time.Millisecond,
				P50:          10 * time.Millisecond,
				P95:          10 * time.Millisecond,
				P99:          10 * time.Millisecond,
			},
		},
		{
			name: "network errors have no status code",
			in: []Result{
				{Error: errors.New("dial failed")},
				{StatusCode: 200, Latency: 15 * time.Millisecond},
			},
			want: Stats{
				Total:        2,
				Successes:    1,
				Errors:       1,
				StatusCounts: map[int]int{200: 1},
				Min:          15 * time.Millisecond,
				Max:          15 * time.Millisecond,
				P50:          15 * time.Millisecond,
				P95:          15 * time.Millisecond,
				P99:          15 * time.Millisecond,
			},
		},
		{
			name: "all requests failed",
			in: []Result{
				{StatusCode: 500},
				{Error: errors.New("timeout")},
			},
			want: Stats{
				Total:        2,
				Successes:    0,
				Errors:       2,
				StatusCounts: map[int]int{500: 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateStats(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("calculateStats() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestPercentile(t *testing.T) {
	latencies := []time.Duration{
		10 * time.Millisecond,
		20 * time.Millisecond,
		30 * time.Millisecond,
		40 * time.Millisecond,
	}

	tests := []struct {
		name string
		p    float64
		want time.Duration
	}{
		{name: "p50", p: 0.50, want: 20 * time.Millisecond},
		{name: "p95", p: 0.95, want: 40 * time.Millisecond},
		{name: "p99", p: 0.99, want: 40 * time.Millisecond},
		{name: "empty", p: 0.95, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := latencies
			if tt.name == "empty" {
				in = nil
			}

			got := percentile(in, tt.p)
			if got != tt.want {
				t.Fatalf("percentile() = %s, want %s", got, tt.want)
			}
		})
	}
}
