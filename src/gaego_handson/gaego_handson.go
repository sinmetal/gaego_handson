package gaego_handson

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/hello", handler)

	m := http.DefaultServeMux
	SetUpItem(m)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
