package main

import (
	"encoding/json"

	"github.com/googollee/go-socket.io"
	l "github.com/sirupsen/logrus"
)

// Function to start socket server
func serveSocket() *socketio.Server {
	server, err := socketio.NewServer(nil)
	if err != nil {
		l.Error(err)
	}

	server.OnConnect("/", func(s socketio.Conn) error {
		l.WithFields(l.Fields{
			"SID": s.ID(),
			"URL": s.URL(),
		}).Info("New connection")
		return nil
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		l.WithFields(l.Fields{
			"SID":   s.ID(),
			"Error": e,
		}).Error("Error in connection")
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		l.WithFields(l.Fields{
			"SID":    s.ID(),
			"Reason": reason,
		}).Warn("Connection lost")
	})

	// Send te state of devices
	server.OnEvent("/", "get_devices", func(s socketio.Conn) {
		jsonState, _ := json.Marshal(deviceState)
		s.Emit("devices_list", string(jsonState))
	})

	// To catch the device selected by the user
	server.OnEvent("/", "select_device", func(s socketio.Conn, pID string) bool {
		l.Info("Device selected by user: ", pID)
		disconnectActivePeripheral()
		go connectPeripheral(pID)
		return true
	})

	return server
}
