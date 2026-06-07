package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	result := fetch(client, "https://www.google.com")
	fmt.Printf("%+v\n", result)
}
