package main

import (
	"context"
	"io"
	"net/http"

	"github.com/googollee/go-socket.io"
	"github.com/markbates/pkger"
	ob "github.com/ramantehlan/pulse/internal/openBrowser"
	s "github.com/ramantehlan/pulse/internal/socket"
	l "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	// Port for the client server.
	Port = ":7000"
)

// Socket state
var socket = s.ServeSocket()

// Function to start the socket
func startSocket() {
	// To catch the device selected by the user
	socket.OnEvent("/", "select_device", func(s socketio.Conn, pID string) bool {
		l.Info("Device selected by user: ", pID)
		connectPeripheral(pID)
		return true
	})

}

// Function to start frontend and websocket
func startServer() {
	go socket.Serve()
	defer socket.Close()

	templateDir := pkger.Dir("/bin/template")
	ob.OpenBrowser("http://localhost" + Port)
	http.Handle("/", http.FileServer(templateDir))
	http.Handle("/socket.io/", socket)

	l.WithFields(l.Fields{"port": Port}).Info("File server running")
	l.Fatal(http.ListenAndServe(Port, nil))
}

func connectPeripheral(pID string) {
	l.Info("Trying to connect to ", pID)

	// create a gRPC stub
	conn, err := grpc.Dial(":7002", grpc.WithInsecure())
	if err != nil {
		l.Error(err)
	}
	defer conn.Close()
	device := NewMibandDeviceClient(conn)
	deviceUUID := &DeviceUUID{
		UUID: pID,
	}

	response, err := device.GetHeartBeats(context.Background(), deviceUUID)
	if err != nil {
		l.Error(err)
	}
	// Fetch stream and broadcast it to the socket
	for {
		pulse, err := response.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			l.Error(err)
			break
		}

		l.Info(pulse)
	}
}

func main() {
	startSocket()
	startServer()
	select {}
}
