package main

import (
	"fmt"
	"gym"
	"log"
	"net/http"
)

func main() {
	r := gym.New()

	r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			log.Printf("[%q]: %q\n", k, v)
			fmt.Fprintf(w, "[%q]: %q\n", k, v)
		}
	})

	r.POST("/post", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Path: %q\n", r.URL.Path)
		fmt.Fprintf(w, "Path: %q\n", r.URL.Path)
	})

	err := r.Run(":9099")
	if err != nil {
		log.Fatalln("Run server ERROR:", err)
	}
}
