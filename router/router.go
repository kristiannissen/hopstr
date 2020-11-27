package router

import (
  "regexp"
  "fmt"
)
// Router
type Router struct {
  Routes map[string]string
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
  fmt.Println(groups[0:])
  return path
}
// Dispatch route to handler
func Dispatch() string {
  return "Hello"
}
