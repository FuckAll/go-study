package main

import (
	"fmt"
	// "log"
	"net/http"
	// "strings"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return

	}
	http.NotFound(w, r)
	return

}
func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9000", mux)

}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello wu")

}
