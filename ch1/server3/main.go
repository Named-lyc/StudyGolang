package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// handler返回http request
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "host = %q\n", r.Host)
	fmt.Fprintf(w, "remotoAddr =%q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "form[%q] = %q\n", k, v)
	}
}
