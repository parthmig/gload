package main

import "fmt"

func main() {
	result := fetch("https://www.google.com")
	fmt.Printf("%+v\n", result)
}
