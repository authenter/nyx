package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

// SetURLVars sets the URL variables for the given request,
// to be accessed via router.Vars for testing route behaviour.
func SetURLVars(r *http.Request, val map[string]string) *http.Request {
	return mux.SetURLVars(r, val)
}

// Vars returns the route variables for the current request, if any.
func Vars(r *http.Request) map[string]string {
	return mux.Vars(r)
}
