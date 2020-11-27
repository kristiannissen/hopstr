package router

import (
	"regexp"
	str "strings"
)

// Router
type Router struct {
	Routes map[string]string
	Params map[string]string
}

// Init function
func init() {
	// fmt.Println("Init called")
}

// Used only to test
func Hello(x int) int {
	return x * 2
}

// Creates a new Router
func NewRouter() *Router {
	return &Router{Routes: make(map[string]string)}
}

// Adds a new route to routes
func (r *Router) AddRoute(path string, s string) {
	var p string = rewrite(path)

	r.Routes[p] = s
}

// Rewrite string to regexp
func rewrite(path string) string {
	reg := regexp.MustCompile(`\{(?P<key>[a-z]+)\}`)
	groups := reg.FindStringSubmatch(path)

	if len(groups) > 0 {
		// fmt.Println("rep is ", groups[0], path)
		// TODO: Add the key to the params map
		// Replace key with regexp string
		path = str.ReplaceAll(path, groups[0], "([a-zA-Z0-9]+)")
	}
	path = str.ReplaceAll(path, "/", "\\/")
	return path
}
