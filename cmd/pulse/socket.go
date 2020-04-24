package main

import (
	"github.com/googollee/go-socket.io"
	l "github.com/sirupsen/logrus"
)

var socket = serveSocket()

// Device is to hold properties of a device
type Device struct {
	name string
	id   string
}

var m = map[int]Device{
	0: Device{
		"mi", "123",
	},
	1: Device{
		"fitbit", "345",
	},
}

// Function to start socket server
func serveSocket() *socketio.Server {
	server, err := socketio.NewServer(nil)
	if err != nil {
		l.Error(err)
	}

	server.OnConnect("/", func(s socketio.Conn) error {
		l.WithFields(l.Fields{
			"SID": s.ID(),
		}).Info("New connection")
		s.Join("broadcastDevices")
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

	// To catch the device selected by the user
	server.OnEvent("/", "select_device", func(s socketio.Conn, msg string) string {
		l.Info("Device selected by user: ", msg)
		return msg
	})

	// To emit the available devices

	return server
}
