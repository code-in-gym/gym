package gym

import (
	"fmt"
	"log"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (e *Engine) addRoute(method, path string, handler HandlerFunc) {
	key := method + path
	e.router[key] = handler
}

func (e *Engine) GET(path string, handler HandlerFunc) {
	e.addRoute(GET, path, handler)
}

func (e *Engine) POST(path string, handler HandlerFunc) {
	e.addRoute(POST, path, handler)
}

func (e *Engine) PUT(path string, handler HandlerFunc) {
	e.addRoute(PUT, path, handler)
}

func (e *Engine) DELETE(path string, handler HandlerFunc) {
	e.addRoute(DELETE, path, handler)
}

func (e *Engine) Run(addr string) error {
	log.Println("Server listening at ", addr)
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, ok := e.router[r.Method+r.URL.Path]; ok {
		handler(w, r)
	} else {
		log.Printf("404 Not Found\n")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 Not Found %q %q\n", r.Method, r.URL.Path)
	}
}
