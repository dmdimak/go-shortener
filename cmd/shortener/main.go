package main

import (
	"io"
	"net/http"
)

var storedUrl string = ""

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, treatUrl)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}

func treatUrl(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error request", http.StatusBadRequest)
			return
		}

		body, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Error parse request", http.StatusBadRequest)
			return
		}

		storedUrl = string(body)

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("http://localhost:8080/EwHXdJfB"))
		return
	} else if r.Method == http.MethodGet {

		w.Header().Set("Location", storedUrl)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	http.Error(w, "Only POST or GET method is allowed", http.StatusBadRequest)
}
