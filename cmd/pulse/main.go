package main

import (
	"net/http"

	"github.com/googollee/go-socket.io"
	"github.com/markbates/pkger"
	"github.com/paypal/gatt"
	"github.com/ramantehlan/pulse/internal"
	l "github.com/sirupsen/logrus"
)

const (
	// Port for the client server.
	Port = ":7000"
)

// DeviceStruct is structure of information that is sent.
type DeviceStruct struct {
	Name             string
	LocalName        string
	PeripheralID     string
	TXPowerLevel     int
	ManufacturerData []byte
	ServiceData      []gatt.ServiceData
}

var deviceState map[string]DeviceStruct
var socket = serveSocket()

// Function to start frontend and websocket
func startServer() {
	go socket.Serve()
	defer socket.Close()

	templateDir := pkger.Dir("/bin/template")

	go internal.OpenBrowser("http://localhost" + Port)
	http.Handle("/", http.FileServer(templateDir))
	http.Handle("/socket.io/", socket)
	l.Fatal(http.ListenAndServe(Port, nil))

	l.WithFields(l.Fields{
		"port": Port,
	}).Info("File server running")
}

func main() {
	go startServer()
	deviceState = make(map[string]DeviceStruct)
	device := searchDevices(onDeviceDiscovered, onDeviceStateChanged)

	// To catch the device selected by the user
	socket.OnEvent("/", "select_device", func(s socketio.Conn, msg string) bool {
		l.Info("Device selected by user: ", msg)
		device.StopScanning()
		return true
	})

	select {}
}
