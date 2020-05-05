package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/bettercap/gatt"
	"github.com/markbates/pkger"
	"github.com/ramantehlan/pulse/api"
	internal "github.com/ramantehlan/pulse/internal/openBrowser"
	options "github.com/ramantehlan/pulse/internal/options"
	l "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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
// Stores the map of type string to store available peripheral
var deviceState = make(map[string]DeviceStruct)

// Peripheral currently connected to
var activePeripheral string = ""
var peripheralState = make(map[string]gatt.Peripheral)
var advertisementState = make(map[string]*gatt.Advertisement)
var socket = serveSocket()

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
	}
}

func connectPeripheral(pID string) {
	l.Info("Trying to connect to ", pID)
	activePeripheral = pID
	selectedPeripheral := peripheralState[pID]
	selectedPeripheral.Device().Connect(selectedPeripheral)
}

func main() {
	// Create a listener on tcp port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		l.Error(err)
	}

	s := api.Server{}
	grpcServer := grpc.NewServer()

	api.RegisterPingServer(grpcServer, &s)

	go startServer()

	d, err := gatt.NewDevice(options.DefaultClientOptions...)
	if err != nil {
		l.Error("Failed to open device, err: %s\n", err)
		return
	}

	d.Handle(
		gatt.PeripheralDiscovered(onPeripheralDiscovered),
		gatt.PeripheralConnected(onPeripheralConnected),
		gatt.PeripheralDisconnected(onPeripheralDisconnected),
	)
	d.Init(onDeviceStateChanged)

	time.Sleep(10 * time.Second)
	d.StopScanning()
	l.Info("Stopping device scan")

	select {}
}
