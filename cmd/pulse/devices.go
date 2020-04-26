package main

import (
	"fmt"
	"log"

	"github.com/paypal/gatt"
	"github.com/paypal/gatt/examples/option"
	l "github.com/sirupsen/logrus"
)

func onDeviceDiscovered(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {
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

// Function to search for BLE devices
func searchDevices(onDiscovered func(gatt.Peripheral, *gatt.Advertisement, int), stateChanged func(gatt.Device, gatt.State)) gatt.Device {
	d, err := gatt.NewDevice(option.DefaultClientOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s\n", err)
		return d
	}

	d.Handle(gatt.PeripheralDiscovered(onDiscovered))
	d.Init(stateChanged)
	return d
}
