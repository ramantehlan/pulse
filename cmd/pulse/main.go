package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/bettercap/gatt"
	"github.com/markbates/pkger"
	heartBeat "github.com/ramantehlan/pulse/api"
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

func startPulseTool() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":7002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := heartBeat.NewDevicePulseClient(conn)

	response, err := c.GetHeartBeats(context.Background(),
		&heartBeat.DeviceUUID{UUID: "uuid"})
	if err != nil {
		l.Error("error: ", err)
	}
	l.Info("Response from gRPC server: %s", response.HeartBeats)
}

func main() {
	startPulseTool()
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
