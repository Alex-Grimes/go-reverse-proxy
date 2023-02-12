package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	reverseProxy := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Printf("[reverse proxy server] recived request at: %s\n", time.Now())

	})

	log.Fatal(http.ListenAndServe(":8080", reverseProxy))
}
