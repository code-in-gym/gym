package gym

import (
	"fmt"
	"log"
	"net/http"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method, path string, handler HandlerFunc) {
	key := method + path
	r.handlers[key] = handler
	log.Printf("[GYM] %6s - %s\n", method, path)
}

func (r *router) handle(c *Context) {
	if handler, ok := r.handlers[c.Method+c.Path]; ok {
		handler(c)
	} else {
		log.Printf("404 Not Found\n")
		c.Writer.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(c.Writer, "404 Not Found %q %q\n", c.Method, c.Path)
	}
}
