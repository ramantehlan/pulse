package main

import (
	"fmt"

	"github.com/paypal/gatt"
)

const (
	// Port for the client server.
	Port = ":8000"
)

// Function to handle when devices are discovered
func onDiscovered(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {

	fmt.Printf("\nPeripheral ID:%s, NAME:(%s)\n", p.ID(), p.Name())
	fmt.Println("  Local Name        =", a.LocalName)
	fmt.Println("  TX Power Level    =", a.TxPowerLevel)
	fmt.Println("  Manufacturer Data =", a.ManufacturerData)
	fmt.Println("  Service Data      =", a.ServiceData)
}

// Function to handle when the devices state is changed
func onStateChanged(d gatt.Device, s gatt.State) {
	fmt.Println("State:", s)
	switch s {
	case gatt.StatePoweredOn:
		fmt.Println("scanning...")
		d.Scan([]gatt.UUID{}, false)
		return
	default:
		d.StopScanning()
	}
}

func main() {

	startClient()
	//searchDevices(onDiscovered, onStateChanged)
}
