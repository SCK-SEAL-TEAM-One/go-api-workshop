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
	fmt.Fprintln(w, "hello world")
}
