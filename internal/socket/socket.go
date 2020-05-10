package socket

import (
	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	l "github.com/sirupsen/logrus"
)

// ServeSocket is to get the base instance of a socket server
func ServeSocket() *gosocketio.Server {
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		l.WithFields(l.Fields{
			"SID": c.Id(),
		}).Info("New connection")

		c.Join("pulse")
	})

	server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		c.Leave("pulse")
		l.WithFields(l.Fields{
			"SID": c.Id(),
		}).Info("Connection dropped")
	})

	//error catching handler
	server.On(gosocketio.OnError, func(c *gosocketio.Channel) {
		l.Error("Error in socket connection")
	})

	return server
}
