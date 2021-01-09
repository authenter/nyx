package router

import (
	"github.com/gorilla/mux"
)

// Router registers routes to be matched and dispatches a handler
type Router struct {
	*mux.Router
}

// New returns a Router objects filled with new gorilla/mux Router
func New() *Router {
	return &Router{mux.NewRouter()}
}
