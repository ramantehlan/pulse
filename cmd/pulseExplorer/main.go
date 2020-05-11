package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/bettercap/gatt"
	explore "github.com/ramantehlan/pulse/internal/exploreDevices"
	"github.com/ramantehlan/pulse/internal/options"
	l "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	// Port for the client server.
	Port = 7004
	// SearchTime is the time limit to search devices for
	SearchTime = 8
)

// create a server instance
var s = explore.DevicesServer{
	DeviceState: make(map[string]explore.DeviceStruct),
}

func onPeripheralDiscovered(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {
	state := explore.DeviceStruct{
		p.Name(),
		a.LocalName,
		p.ID(),
		a.TxPowerLevel,
		a.ManufacturerData,
		a.ServiceData,
	}

	s.AddDeviceState(state)

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

func searchDevices() {
	d, err := gatt.NewDevice(options.DefaultClientOptions...)
	if err != nil {
		l.Error("Failed to open device, err: ", err)
		return
	}
	d.Handle(
		gatt.PeripheralDiscovered(onPeripheralDiscovered),
	)
	d.Init(onDeviceStateChanged)
}

func startServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", Port))
	fmt.Println("pulseExplorer listening on ", Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a gRPC server object
	grpcServer := grpc.NewServer()
	explore.RegisterExploreDevicesServer(grpcServer, &s)
	grpcServer.Serve(lis)
}

// We intentionally want this program to exit after x seconds
// as this program controls the bluetooth, so when it's running
// we can't connect to miband or other tools
func main() {
	go searchDevices()
	go startServer()
	time.Sleep(SearchTime * time.Second)
}
