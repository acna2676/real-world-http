package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://localhost:18443")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	dump, err := http.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(dump))
}
