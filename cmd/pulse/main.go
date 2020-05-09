package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/googollee/go-socket.io"
	"github.com/markbates/pkger"
	explore "github.com/ramantehlan/pulse/internal/exploreDevices"
	ob "github.com/ramantehlan/pulse/internal/openBrowser"
	s "github.com/ramantehlan/pulse/internal/socket"
	l "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	// Port for the client server.
	Port = ":7000"
)

// Socket state
var socket = s.ServeSocket()

// pulseState
var pulseState = make(map[string][]int)

// Device state
var deviceState = make(map[string]explore.DeviceStruct)

// Function to start the socket
func startSocket() {
	// To catch the device selected by the user
	socket.OnEvent("/", "select_device", func(s socketio.Conn, pID string) bool {
		l.Info("Device selected by user: ", pID)
		connectPeripheral(pID)
		return true
	})

	/**
	Ideally, we should be broadcasting to the room, but since that feature is not working
	for me right now, so using the (wrong)temp method

	Maybe we should use golang-socketio
	**/
	socket.OnEvent("/", "get_devices", func(s socketio.Conn) bool {
		jsonState, _ := json.Marshal(deviceState)
		s.Emit("devices_list", string(jsonState))
		return true
	})

	socket.OnEvent("/", "get_pulse", func(s socketio.Conn) bool {
		jsonState, _ := json.Marshal(pulseState)
		s.Emit("heartBeat", string(jsonState))
		return true
	})
}

// Function to start frontend and websocket
func startServer() {
	go socket.Serve()
	defer socket.Close()

	templateDir := pkger.Dir("/bin/template")
	ob.OpenBrowser("http://localhost" + Port)
	http.Handle("/", http.FileServer(templateDir))
	http.Handle("/socket.io/", socket)

	l.WithFields(l.Fields{"port": Port}).Info("File server running")
	l.Fatal(http.ListenAndServe(Port, nil))
}

func getDevices() {
	l.Info("connecting to pulseExplore on port 7004")
	// create a gRPC stub
	conn, err := grpc.Dial(":7004", grpc.WithInsecure())
	if err != nil {
		l.Error(err)
	}
	defer conn.Close()

	client := explore.NewExploreDevicesClient(conn)

	empty := &explore.Empty{}
	stream, err := client.GetList(context.Background(), empty)
	if err != nil {
		l.Error(err)
	}

	for {
		device, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			l.Error(err)
		}
		l.Info(device)

		d := explore.DeviceStruct{
			PeripheralID: device.PID,
			Name:         device.Name,
		}

		deviceState[device.PID] = d
	}

}

func connectPeripheral(pID string) {
	l.Info("Trying to connect to ", pID)
	l.Info("connecting to mibandPulse on port 7002")

	// Reset the value of pulseState
	pulseState["pulse"] = []int{}

	// create a gRPC stub
	conn, err := grpc.Dial(":7002", grpc.WithInsecure())
	if err != nil {
		l.Error(err)
	}
	defer conn.Close()
	device := NewMibandDeviceClient(conn)
	deviceUUID := &DeviceUUID{
		UUID: pID,
	}

	response, err := device.GetHeartBeats(context.Background(), deviceUUID)
	if err != nil {
		l.Error(err)
	}
	// Fetch stream and broadcast it to the socket
	for {
		pulse, err := response.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			l.Error(err)
			break
		}

		l.Info(pulse)
		p, _ := strconv.Atoi(pulse.Pulse)
		pulseState["pulse"] = append(pulseState["pulse"], p)
	}
}

func main() {
	startSocket()
	getDevices()
	startServer()
	select {}
}
