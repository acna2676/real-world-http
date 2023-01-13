package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "Hello from origin server")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("origin servert start at :9001")
	log.Fatalln(http.ListenAndServe(":9001", nil))
}
