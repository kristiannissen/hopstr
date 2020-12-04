package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	str "strings"
)

type Handle func(http.ResponseWriter, *http.Request)

type Route struct {
	routes map[string]Handle
}

func NewRoute() *Route {
	return &Route{routes: make(map[string]Handle)}
}

func (route *Route) HandleFunc(
	pattern string, f func(http.ResponseWriter, *http.Request)) {
	var regPath string = rewritePath(pattern)
	route.routes[regPath] = f
}

func (route *Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for p, h := range route.routes {
		match, _ := regexp.MatchString(p, r.URL.Path)

		if match == true {
			h(w, r)
			return
		}
	}
	http.NotFound(w, r)
}

func rewritePath(path string) string {
	reg := regexp.MustCompile(`\{(?P<key>[a-z]+)}`)
	groups := reg.FindStringSubmatch(path)

	if len(groups) > 0 {
		path = str.ReplaceAll(path, groups[0], "([a-zA-Z0-9]+)")
	}
	path = str.ReplaceAll(path, "/", "\\/")
	return path
}

func HelloKitty(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Kitty path %s method %s", r.URL.Path, r.Method)
}

func main() {
	var port string = os.Getenv("PORT")

	if port == "" {
		port = "80"
	}

	route := NewRoute()
	route.HandleFunc("/", HelloKitty)
	route.HandleFunc("/hello/", HelloKitty)
	route.HandleFunc("/hello/{pussy}/", HelloKitty)

	log.Fatal(http.ListenAndServe(":"+port, route))
}
