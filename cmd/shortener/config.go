package main

import "flag"

var flagRunAddr string
var baseShortenedURL string

var ServerAddress = "localhost:8080"
var BaseUrl = "http://localhost:8080"

func parseFlags() {
	flag.StringVar(&flagRunAddr, "a", ServerAddress, "address to run server")
	flag.StringVar(&baseShortenedURL, "b", BaseUrl, "base address for shortened URL")
	flag.Parse()
}
