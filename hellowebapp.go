package main

import (
	"net/http"
	"fmt"
)

var version = "undefined"

func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Healthy")
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, 
		"<h1>Fancy Hello World Web App</h1>" +
		"<h2>Environment: <span style='background-color: %s'>%s</span></h2>" +
		"<h2>Version: <span>%s</span></h2>",
		"lime", "PRD", version)
}

func main() {
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":80", nil)
}
