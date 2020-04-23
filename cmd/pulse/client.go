package pulse

import (
	"log"
	"net/http"

	"github.com/markbates/pkger"
)

// Function to start the client server.
func startClient() {
	go socket.Serve()
	defer socket.Close()

	templateDir := pkger.Dir("/bin/template")

	http.Handle("/socket.io/", socket)
	http.Handle("/", http.FileServer(templateDir))
	log.Fatal(http.ListenAndServe(Port, nil))
}
