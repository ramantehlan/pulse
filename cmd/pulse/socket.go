package main

import (
	"github.com/googollee/go-socket.io"
	l "github.com/sirupsen/logrus"
)

/**
* Ideally we should be using server.BroadcastToRoom.
* But that is not working, so we are using this weird way.
* Where once the connection is created, we store its instance
* in a global variable, which is not ideal, but works for now.
**/
var socketInstance socketio.Conn

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
		s.SetContext("")
		l.WithFields(l.Fields{
			"SID": s.ID(),
			"URL": s.URL(),
		}).Info("New connection")
		socketInstance = s
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
		s.SetContext(msg)
		l.Info("Device selected by user: ", msg)
		return msg
	})

	return server
}
