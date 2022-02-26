package gym

import (
	"log"
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) GET(path string, handler HandlerFunc) {
	e.router.addRoute(GET, path, handler)
}

func (e *Engine) POST(path string, handler HandlerFunc) {
	e.router.addRoute(POST, path, handler)
}

func (e *Engine) PUT(path string, handler HandlerFunc) {
	e.router.addRoute(PUT, path, handler)
}

func (e *Engine) DELETE(path string, handler HandlerFunc) {
	e.router.addRoute(DELETE, path, handler)
}

func (e *Engine) Run(addr string) error {
	log.Println("Server listening at ", addr)
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.router.handle(NewContext(w, r))
}
