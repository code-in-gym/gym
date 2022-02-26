package main

import (
	"fmt"
	"gym"
	"log"
	"net/http"
)

func main() {
	r := gym.New()

	r.GET("/hello", func(c *gym.Context) {
		for k, v := range c.Req.Header {
			log.Printf("[%q]: %q\n", k, v)
			fmt.Fprintf(c.Writer, "[%q]: %q\n", k, v)
		}
	})

	r.POST("/post", func(c *gym.Context) {
		log.Printf("Path: %q\n", c.Path)
		c.JSON(http.StatusOK, gym.H{
			"path": c.Path,
			"form": c.PostFrom("val"),
		})
	})

	err := r.Run(":9099")
	if err != nil {
		log.Fatalln("Run server ERROR:", err)
	}
}
