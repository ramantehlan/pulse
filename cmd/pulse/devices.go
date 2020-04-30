package main

import (
	"fmt"

	"github.com/bettercap/gatt"
	mi "github.com/ramantehlan/pulse/pkg/miband"
	l "github.com/sirupsen/logrus"
)

func onPeripheralDiscovered(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {
	peripheralState[p.ID()] = p
	advertisementState[p.ID()] = a
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

func onPeripheralConnected(p gatt.Peripheral, err error) {
	l.Info("Peripheral Connected: ", p.ID())
	l.Info("Received Signal Strength Indicator (RSSI): ", p.ReadRSSI())

	fmt.Println(mi.UUIDServiceMiband2)
	genericAccessUUID := []gatt.UUID{gatt.UUID16(mi.UUIDServiceGenericAccess)}
	fmt.Println(genericAccessUUID)
	ss, err := p.DiscoverServices(genericAccessUUID)
	fmt.Println(len(ss))
	if err != nil {
		fmt.Println(" Error in discovering services: ", err)
	}
}

func onPeripheralDisconnected(p gatt.Peripheral, err error) {
	l.Info("Peripheral Disconnected: ", p.ID())
}
