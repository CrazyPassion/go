package helloworld

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloNameMux(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayhelloNameMux(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")
}

func MymuxTest() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
