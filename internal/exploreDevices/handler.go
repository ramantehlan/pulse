package exploreDevices

import (
	"log"

	"github.com/bettercap/gatt"
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

// DevicesServer represets the gRPC server
type DevicesServer struct {
	DeviceState map[string]DeviceStruct
}

// GetList is to handle when the request is sent server
func (s *DevicesServer) GetList(empty *Empty, stream ExploreDevices_GetListServer) error {
	log.Printf("Device list requested")

	for _, value := range s.DeviceState {
		device := Device{
			PID:  value.PeripheralID,
			Name: value.Name,
		}
		if err := stream.Send(&device); err != nil {
			return err
		}
	}

	return nil
}

// AddDeviceState is to add the device state to the list
func (s *DevicesServer) AddDeviceState(ds DeviceStruct) error {
	s.DeviceState[ds.PeripheralID] = ds
	return nil
}
