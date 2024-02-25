package main

import (
	"fmt"
	helm_chart "helm.sh/helm/v3/pkg/chart"
	"html"
	"log"
	"net/http"
)

func main() {
	v := helm_chart.Chart{}
	v = v
	//fmt.Println(metadata.SPECIFICATION_VERSION)
	// comment
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi There Yolanda!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
