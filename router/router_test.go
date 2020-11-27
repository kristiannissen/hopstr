package router

import (
  "testing"
  "fmt"
  r "router"
)

func TestHello(t *testing.T) {
  x := r.Hello(2)
  fmt.Println("Hello Kitty")

  if x != 4 {
    t.Error("Expected 2 * 2 to equal 4")
  }
}

func TestNewRoute(t *testing.T) {
  route := r.NewRouter()

  if len(route.Routes) > 0 {
    t.Errorf("New is not returning empty map")
  }
}

func TestAddRoute(t *testing.T) {
  route := r.NewRouter()
  route.AddRoute("kitty", "Kitty")
  route.AddRoute("pussy", "Pussy")

  if len(route.Routes) != 2 {
    t.Error("AddRoute did not add any routes")
  }
}

func TestAddRouteRewrite(t *testing.T) {
  route := r.NewRouter()
  route.AddRoute("/hello/{pussy}/", "hello")
  route.AddRoute("/{kitty}/hello/{pussy}/", "hello")

  t.Error("shit")
}
