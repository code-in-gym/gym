package main

import (
	"gym"
	"log"
	"net/http"
)

func main() {
	r := new(gym.Engine)
	err := http.ListenAndServe(":9099", r)
	if err != nil {
		log.Fatalln("Run server ERROR:", err)
	}
}
