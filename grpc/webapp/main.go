package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func main() {

	x := 1

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name, _ := os.Hostname()
		fmt.Fprintf(w, "HostName %q", name)
		fmt.Fprintf(w, "Hello %q", html.EscapeString(r.URL.Path))

	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		x = x + 1
		fmt.Fprintf(w, "healthprobe %v ", x)
		if x%5 == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "unhealthy")
		} else {
			w.WriteHeader(http.StatusAccepted)
			fmt.Fprintf(w, "yaayy healthy")
		}

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
