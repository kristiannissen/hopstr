package router

import (
  "testing"
  "fmt"
  r "router"
)

func TestHello(t *testing.T) {
  x := r.Hello(2)
  fmt.Print(x)
  if x != 4 {
    t.Error("Expected 2 * 2 to equal 4")
  }
}
