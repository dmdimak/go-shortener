package main

import "flag"

var flagRunAddr string
var baseShortenedURL string

func parseFlags() {
	flag.StringVar(&flagRunAddr, "a", "localhost:8080", "address to run server")
	flag.StringVar(&baseShortenedURL, "b", "http://localhost:8080", "base address for shortened URL")
	flag.Parse()
}
