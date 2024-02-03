package main

import "flag"

var flagRunAddr string
var baseShortenedURL string

var ServerAddress = "localhost:8080"
var BaseURL = "http://localhost:8080" // fix this var name

func parseFlags() {
	flag.StringVar(&flagRunAddr, "a", ServerAddress, "address to run server")
	flag.StringVar(&baseShortenedURL, "b", BaseURL, "base address for shortened URL")
	flag.Parse()
}
