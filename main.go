package main

import (
	"fmt"
	"net/http"
)

func main() {

	config := parseConfig()

	if err := validateConfig(config); err != nil {
		fmt.Println("Error:", err)
		return
	}

	client := &http.Client{
		Timeout: config.Timeout,
	}

	results := runLoadTest(client, config)

	summary := calculateStats(results)
	printSummary(config, summary)

}
