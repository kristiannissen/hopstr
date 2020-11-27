package dispatch

import (
	d "dispatch"
	"fmt"
	"net/http"
	"testing"
)

func TestAddRoute(t *testing.T) {
	disp := d.NewRoute()
	disp.AddRoute("/hello/{kitty}/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello from the other side")
	})

	fmt.Println(disp)
}
