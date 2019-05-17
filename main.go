package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// used to dump headers for debugging
func handler(w http.ResponseWriter, r *http.Request) {

	// disable cache
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	// set hostname (used for demo)
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprint(w, "Error:", err)
	}
	fmt.Fprintf(w, "Hello from %v\n", hostname)

	return

}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":5005", nil))
}
