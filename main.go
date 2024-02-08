package main

import (
	"fmt"
	"github.com/theupdateframework/go-tuf/v2/metadata"
	"html"
	"log"
	"net/http"
)

func main() {
	fmt.Println(metadata.SPECIFICATION_VERSION)
	// comment
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi There Yolanda!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
