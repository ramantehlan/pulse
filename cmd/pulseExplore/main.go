package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/bettercap/gatt"
	"github.com/googollee/go-socket.io"
	api "github.com/ramantehlan/pulse/internal/exploreDevices"
	"github.com/ramantehlan/pulse/internal/options"
	s "github.com/ramantehlan/pulse/internal/socket"
	l "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	// Port for the client server.
	Port = ":7004"
	// SearchTime is the time limit to search devices for
	SearchTime = 8000
)

// Socket state
var socket = s.ServeSocket()

// Global variables to manage state
// Stores the map of type string to store available peripheral
var deviceState = make(map[string]DeviceStruct)

// DeviceStruct is structure of information that is sent.
type DeviceStruct struct {
	Name             string
	LocalName        string
	PeripheralID     string
	TXPowerLevel     int
	ManufacturerData []byte
	ServiceData      []gatt.ServiceData
}

// Function to start the socket
func startSocket() {

	// Not an ideal method to ask client to emit to get_devices first
	// and the emit the devices_list, but the socket.io library is not
	// working for the broadcasting.
	// Send te state of devices
	socket.OnEvent("/", "get_devices", func(s socketio.Conn) {
		jsonState, _ := json.Marshal(deviceState)
		s.Emit("devices_list", string(jsonState))
	})

	go socket.Serve()
	defer socket.Close()

	http.Handle("/socket.io/", socket)
	l.WithFields(l.Fields{"port": Port}).Info("pulseExplorer server running")
	l.Fatal(http.ListenAndServe(Port, nil))
}

func onPeripheralDiscovered(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {
	deviceState[p.ID()] = DeviceStruct{
		p.Name(),
		a.LocalName,
		p.ID(),
		a.TxPowerLevel,
		a.ManufacturerData,
		a.ServiceData,
	}

	l.WithFields(l.Fields{
		"Name":         p.Name(),
		"LocalName":    a.LocalName,
		"PeripheralID": p.ID(),
	}).Info("New device discovered")
}

func onDeviceStateChanged(d gatt.Device, s gatt.State) {
	l.Info("BT Device Status:", s)
	switch s {
	case gatt.StatePoweredOn:
		l.Info("Starting to scan device...")
		d.Scan([]gatt.UUID{}, false)
		return
	default:
		l.Error("Permission denied or device not found")
		d.StopScanning()
	}
}

func sendRequest(vhandle uint16, b []byte, per gatt.Peripheral) error {
	c := &gatt.Characteristic{}
	c.SetVHandle(vhandle)
	return per.WriteCharacteristic(c, b, true)
}

// DevicesServer represets the gRPC server
type DevicesServer struct {
}

// GetList is to handle when the request is sent server
func (s *DevicesServer) GetList(empty *api.Empty, stream api.ExploreDevices_GetListServer) error {
	log.Printf("Device list requested")

	for _, value := range deviceState {

		device := api.Device{
			PID:  value.PeripheralID,
			Name: value.Name,
		}

		if err := stream.Send(&device); err != nil {
			return err
		}

	}

	return nil
}

func main() {
	go startSocket()

	d, err := gatt.NewDevice(options.DefaultClientOptions...)
	if err != nil {
		l.Error("Failed to open device, err: ", err)
		return
	}

	d.Handle(
		gatt.PeripheralDiscovered(onPeripheralDiscovered),
	)
	d.Init(onDeviceStateChanged)
	d.StopScanning()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7005))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	s := DevicesServer{}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	api.RegisterExploreDevicesServer(grpcServer, &s)
	grpcServer.Serve(lis)

	time.Sleep(SearchTime * time.Second)
}
