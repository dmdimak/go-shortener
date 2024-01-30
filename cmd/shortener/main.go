package main

import (
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

var storedURL string = ""

// update

func main() {

	parseFlags()

	r := chi.NewRouter()
	r.MethodFunc("GET", "/{id}", treatURL)
	r.MethodFunc("POST", "/", treatURL)

	err := http.ListenAndServe(flagRunAddr, r)
	if err != nil {
		panic(err)
	}
}

func treatURL(w http.ResponseWriter, r *http.Request) {

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

		storedURL = string(body)

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(baseShortenedURL + "/EwHXdJfB"))
		return
	} else if r.Method == http.MethodGet {
		w.Header().Set("Location", storedURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	http.Error(w, "Only POST or GET method is allowed", http.StatusBadRequest)
}
