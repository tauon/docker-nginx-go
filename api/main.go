package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	instance = flag.Int("instance", 0, "Instance number")
)

func main() {
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi, from %d.", *instance)
	})

	http.ListenAndServe(":60080", nil)
}
