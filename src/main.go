package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestBody struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

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
		query := r.URL.Query()
		text := query.Get("text")
		if text != "" {
			fmt.Fprint(w, "hello ", text)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "don't use post method")
	}
}

func HelloWithName(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var requestBody RequestBody
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&requestBody)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, "hello ", requestBody.Name)
	}
}
