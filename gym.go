package gym

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		log.Printf("Path: %q\n", r.URL.Path)
		fmt.Fprintf(w, "Path: %q\n", r.URL.Path)
	case "/hello":
		for k, v := range r.Header {
			log.Printf("[%q]: %q\n", k, v)
			fmt.Fprintf(w, "[%q]: %q\n", k, v)
		}
	default:
		log.Printf("404 Not Found\n")
		fmt.Fprintf(w, "404 Not Found\n")
	}
}
