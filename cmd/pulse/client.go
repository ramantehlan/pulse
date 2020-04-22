package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	// Port for the client server.
	Port = ":8000"
	// ClientPage is to specifiy the index.html to serve.
	ClientPage = "../../client/out/index.html"
)

// Function to load the clientPage
func serveStatic(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(ClientPage)
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}

// Function to start the client server.
func startClient() {
	socket = serveSocket()
	go socket.Serve()
	defer socket.Close()

	http.Handle("/socket.io/", socket)
	http.HandleFunc("/", serveStatic)
	log.Fatal(http.ListenAndServe(Port, nil))
}
