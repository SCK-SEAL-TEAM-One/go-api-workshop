package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloworld)
	http.ListenAndServe(":3000", nil)
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		query := r.URL.Query()
		text := query.Get("text")
		if text != "" {
			fmt.Fprint(w, "hello ", text)
			return
		}
		fmt.Fprint(w, "hello world")
	}
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "don't use post method")
	}
}
