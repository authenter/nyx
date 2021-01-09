package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Route stores information to match a request and build URLs
type Route struct {
	*mux.Route
}

// CurrentRoute returns the matched route for the current request, if any
func CurrentRoute(r *http.Request) *Route {
	return &Route{mux.CurrentRoute(r)}
}
