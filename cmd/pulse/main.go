package main

import (
	"fmt"
	"net/http"
	"time"

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

// Global variables to manage state
var deviceState = make(map[string]DeviceStruct)
var activePeripheral string = ""
var peripheralState = make(map[string]gatt.Peripheral)
var advertisementState = make(map[string]*gatt.Advertisement)
var socket = serveSocket()
var done = make(chan struct{})

// Function to start frontend and websocket
func startServer() {
	go socket.Serve()
	defer socket.Close()

	templateDir := pkger.Dir("/bin/template")
	internal.OpenBrowser("http://localhost" + Port)
	http.Handle("/", http.FileServer(templateDir))
	http.Handle("/socket.io/", socket)

	l.WithFields(l.Fields{"port": Port}).Info("File server running")
	l.Fatal(http.ListenAndServe(Port, nil))
}

func disconnectActivePeripheral() {
	if activePeripheral != "" {
		p := peripheralState[activePeripheral]
		p.Device().CancelConnection(p)
		l.Warn("Disconnecting from ", activePeripheral)
	}
}

// Function to get heartbeats
func connectPeripheral(pID string) {
	l.Info("Trying to connect to ", pID)
	activePeripheral = pID
	selectedPeripheral := peripheralState[pID]
	selectedPeripheral.Device().Connect(selectedPeripheral)
	l.Info("Device Connected")
	_, err := selectedPeripheral.DiscoverServices(nil)
	if err != nil {
		fmt.Println("Failed to discover services, err: %s\n", err)
	}

}

func main() {
	go startServer()
	device := searchDevices(onPeripheralDiscovered, onDeviceStateChanged)

	time.Sleep(10 * time.Second)
	device.StopScanning()
	l.Info("Stopping device scan")

	<-done
	fmt.Println("Done")
}
