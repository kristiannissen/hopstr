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
	route.routes[pattern] = f
}

func (route *Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for p, h := range route.routes {
		var path string = p
		// Does p contain regexp
		reg := regexp.MustCompile(`\{([a-z]+)\}`)
		// Find groups matching
		groups := reg.FindAllStringSubmatch(p, -1)
		// If groups has len > 0
		if len(groups) > 0 {
			for _, v := range groups {
				path = str.ReplaceAll(path, v[0], "([a-zA-Z0-9]+)")
			}
		}
		// Escape / append ^ prepend $
		path = "(?m)^" + str.ReplaceAll(path, "/", "\\/") + "$"
		// Compile full regexp
		var re = regexp.MustCompile(path)
		// Find all submatches
		for _, m := range re.FindAllStringSubmatch(r.URL.Path, -1) {
			// Skip the first match
			fmt.Println(m)
			for i, p := range m[1:] {
				// Insert value by key
				fmt.Println(p, i)
			}
		}
		// Match regular expression path with URL.Path
		m, _ := regexp.MatchString(path, r.URL.Path)
		if m == true {
			h(w, r)
			return
		}
	}
	http.NotFound(w, r)
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
	route.HandleFunc("/hello/{kitty}/eatmy/{pussy}/", HelloKitty)

	log.Fatal(http.ListenAndServe(":"+port, route))
}
