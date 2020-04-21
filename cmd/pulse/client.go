package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const (
	// Port for the client server.
	Port = ":8000"
	// ClientPage is to specifiy the index.html to serve.
	ClientPage = "../../client/out/index.html"
)

// serveStatic is to load the clientPage
func serveStatic(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(ClientPage)
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}

// StartClient is to start the client server.
func startClient() {
	http.HandleFunc("/", serveStatic)
	http.ListenAndServe(Port, nil)
}
