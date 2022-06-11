package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	v1 "github.com/Maximo-Miranda/challenge-fullstack/go-jsonplaceholder-proxy/internal/proxy/v1"
)

func main() {

	port := os.Getenv("PROXY_PORT")
	if len(port) == 0 {
		port = "8080"
	}

	proxyHandler := v1.Proxy{}

	reverseProxy := http.HandlerFunc(proxyHandler.Handler)

	log.Println("Starting proxy server on", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), reverseProxy); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
