package main

import (
	"net/http"

	"github.com/markbates/pkger"
	l "github.com/sirupsen/logrus"
)

// Function to start the client server.
func startClient() {
	go socket.Serve()
	defer socket.Close()

	templateDir := pkger.Dir("/bin/template")
	//templateDir := http.Dir("./bin/template")

	http.Handle("/", http.FileServer(templateDir))
	http.Handle("/socket.io/", socket)
	l.Fatal(http.ListenAndServe(Port, nil))
	l.WithFields(l.Fields{
		"port": Port,
	}).Info("Server Running on localhost")
}
