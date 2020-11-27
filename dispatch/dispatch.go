package dispatch

import (
	"fmt"
	"net/http"
	"regexp"
	str "strings"
)

type Route struct {
	Routes map[string]func(w http.ResponseWriter, r *http.Request)
}

func NewRoute() *Route {
	return &Route{Routes: make(map[string]func(w http.ResponseWriter, r *http.Request))}
}

func (r *Route) AddRoute(str string, handler func(w http.ResponseWriter, r *http.Request)) {
	var uri string = rewrite(str)
	r.Routes[uri] = handler
}

func (r *Route) Dispatch(str string) {
	for k, v := range r.Routes {
		fmt.Printf("Hello %s %s", k, v)
	}
}

func rewrite(path string) string {
	reg := regexp.MustCompile(`\{(?P<key>[a-z]+)\}`)
	groups := reg.FindStringSubmatch(path)

	if len(groups) > 0 {
		// Replace key with regexp string
		path = str.ReplaceAll(path, groups[0], "([a-zA-Z0-9]+)")
	}
	path = str.ReplaceAll(path, "/", "\\/")
	return path
}
