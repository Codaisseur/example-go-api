package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hello)
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s\n", req.Proto, req.URL)
	fmt.Fprintln(w, "<!doctype html><html>")
	fmt.Fprintln(w, "<head><meta charset=\"utf-8\"><title>Ohai! Codaisseur Test App</title></head>")
	fmt.Fprintln(w, "<body><header><h1>Ohai!</h1></header><p>Welcome to Codaisseur!</p><footer>This is just an example app for testing purposes.</footer></body>")
	fmt.Fprintln(w, "</html>")
}
