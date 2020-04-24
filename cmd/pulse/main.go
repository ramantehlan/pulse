package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/markbates/pkger"
	"github.com/paypal/gatt"
	l "github.com/sirupsen/logrus"
)

const (
	// Port for the client server.
	Port = ":7000"
)

// Function to handle when devices are discovered
func onDeviceDiscovered(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {
	if socketInstance == nil {
		l.Debug("Socket Instance Empty")
	} else {
		l.Debug("Socket Started")
		socketInstance.Emit("devices_list", p.ID())
	}

	l.WithFields(l.Fields{
		"Name":              p.Name(),
		"Local Name":        a.LocalName,
		"Peripheral ID":     p.ID(),
		"TX Power Level":    a.TxPowerLevel,
		"Manufacturer Data": a.ManufacturerData,
		"Service Data":      a.ServiceData,
	}).Info("New device discovered")
}

// Function to handle when the devices state is changed
func onDeviceStateChanged(d gatt.Device, s gatt.State) {
	fmt.Println("State:", s)
	switch s {
	case gatt.StatePoweredOn:
		fmt.Println("scanning...")
		d.Scan([]gatt.UUID{}, false)
		return
	default:
		d.StopScanning()
	}
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

// Function to start frontend and websocket
func startServer() {
	socket := serveSocket()
	go socket.Serve()
	defer socket.Close()

	templateDir := pkger.Dir("/bin/template")

	go openBrowser("http://localhost" + Port)
	http.Handle("/", http.FileServer(templateDir))
	http.Handle("/socket.io/", socket)
	l.Fatal(http.ListenAndServe(Port, nil))

	l.WithFields(l.Fields{
		"port": Port,
	}).Info("File server running")

}

func main() {
	startServer()
	//searchDevices(onDeviceDiscovered, onDeviceStateChanged)
}
