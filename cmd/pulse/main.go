package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/graarh/golang-socketio"
	"github.com/markbates/pkger"
	explore "github.com/ramantehlan/pulse/internal/exploreDevices"
	miPulse "github.com/ramantehlan/pulse/internal/mibandDevice"
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

// Device state
var deviceState = make(map[string]explore.DeviceStruct)

// SelectMessage struct for socketio
type SelectMessage struct {
	PID string `json:"pid"`
}

// Function to initialize the socket
func initializeSocket() {
	// To catch the device selected by the user
	socket.On("select_device", func(c *gosocketio.Channel, msg SelectMessage) bool {
		l.Info("Device selected by user: ", msg.PID)
		connectPeripheral(msg.PID)
		return true
	})

	socket.On("get_devices", func(c *gosocketio.Channel) bool {
		jsonState, _ := json.Marshal(deviceState)
		c.Emit("devices_list", string(jsonState))
		return true
	})
}

// Function to start frontend and websocket
func startServer() {
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

	// create a gRPC stub
	conn, err := grpc.Dial(":7002", grpc.WithInsecure())
	if err != nil {
		l.Error("Error in gRPC dial :7002 > ", err)
	}
	defer conn.Close()
	l.Info("Successfully connected to mibandPulse gRPC on 7002 :)")
	device := miPulse.NewMibandDeviceClient(conn)
	deviceUUID := &miPulse.DeviceUUID{
		UUID: pID,
	}

	response, err := device.GetHeartBeats(context.Background(), deviceUUID)
	if err != nil {
		l.Error("Error in calling GetHeartBeats > ", err)
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
		socket.BroadcastTo("pulse", "heartBeat", p)
	}
}

func main() {
	initializeSocket()
	l.Info("Server Starting in 5 seonds | Waiting for pulseExplore to find devices")
	time.Sleep(5 * time.Second)
	getDevices()
	startServer()
	select {}
}
