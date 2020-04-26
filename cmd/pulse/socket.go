package main

import (
	"encoding/json"
	"time"

	"github.com/googollee/go-socket.io"
	l "github.com/sirupsen/logrus"
)

/**
* Ideally we should be using server.BroadcastToRoom.
* But that is not working, so we are using this below logic.
* Where we keep emiting the state of discovered devices in loop.
**/
func emitDevices(s socketio.Conn) {
	for true {
		jsonState, _ := json.Marshal(deviceState)
		s.Emit("devices_list", string(jsonState))
		time.Sleep(5 * time.Second)
	}
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
			"URL": s.URL(),
		}).Info("New connection")
		go emitDevices(s)
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
	server.OnEvent("/", "select_device", func(s socketio.Conn, msg string) bool {
		l.Info("Device selected by user: ", msg)
		return true
	})

	return server
}
