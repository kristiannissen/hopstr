package router

// Router
type Router struct {
  routes map[string]string
}
// Used only to test
func Hello(x int) int {
  return x * 2
}

func (r *Router) AddRoute(path string, s string) *Router {
  r.routes[path] = s
  return r
}
