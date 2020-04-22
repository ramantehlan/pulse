package main

import (
	"log"

	"github.com/paypal/gatt"
	"github.com/paypal/gatt/examples/option"
)

// Function to search for BLE devices
func searchDevices(onDiscovered func(gatt.Peripheral, *gatt.Advertisement, int), stateChanged func(gatt.Device, gatt.State)) {
	d, err := gatt.NewDevice(option.DefaultClientOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s\n", err)
		return
	}

	d.Handle(gatt.PeripheralDiscovered(onDiscovered))
	d.Init(stateChanged)
	select {}
}
