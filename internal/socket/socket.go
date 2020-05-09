package socket

import (
	"github.com/googollee/go-socket.io"
	l "github.com/sirupsen/logrus"
)

// ServeSocket is to get the base instance of a socket server
func ServeSocket() *socketio.Server {
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

	return server
}
