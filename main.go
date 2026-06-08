package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	config := parseConfig()

	if err := validateConfig(config); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	client := &http.Client{
		Timeout: config.Timeout,
	}

	results := runLoadTest(client, config)

	summary := calculateStats(results)
	printSummary(config, summary)

}
