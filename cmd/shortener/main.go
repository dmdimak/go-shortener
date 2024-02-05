package main

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

var storedURL string = ""
var sugar zap.SugaredLogger

type (
	responseData struct {
		status int
		size   int
	}

	loggingResponseWriter struct {
		http.ResponseWriter
		responseData *responseData
	}
)

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size
	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode
}

// update

func main() {

	parseFlags()

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	defer logger.Sync()
	sugar = *logger.Sugar()

	r := chi.NewRouter()
	r.MethodFunc("GET", "/{id}", treatURL)
	r.MethodFunc("POST", "/", treatURL)
	r.Use(withLogger)

	if runAddr := os.Getenv("SERVER_ADDRESS"); runAddr != "" {
		flagRunAddr = runAddr
	}

	if baseURL := os.Getenv("BASE_URL"); baseURL != "" {
		baseShortenedURL = baseURL
	}

	sugar.Infow(
		"Starting server",
		"addr", flagRunAddr,
		"baseURL", baseShortenedURL,
	)

	err = http.ListenAndServe(flagRunAddr, r)
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

func withLogger(h http.Handler) http.Handler {
	logFn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		responseData := &responseData{
			status: 0,
			size:   0,
		}

		lw := loggingResponseWriter{
			ResponseWriter: w,
			responseData:   responseData,
		}

		uri := r.RequestURI
		method := r.Method

		h.ServeHTTP(&lw, r)

		duration := time.Since(start)

		sugar.Infoln(
			"uri", uri,
			"method", method,
			"status", responseData.status,
			"duration", duration,
			"size", responseData.size,
		)
	}

	return http.HandlerFunc(logFn)
}
