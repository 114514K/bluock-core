package handler

import (
	"fmt"
	"log"

	"tinygo.org/x/bluetooth"
)

const (
	loggingTemplate = "Address: %s RSSI: %d%s"
	deviceNameTemplate = " Name: %s"
)

type ConnectionTestHandler struct {}

func (handler ConnectionTestHandler) OnScan(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
	deviceName := device.LocalName()
	deviceNameMessage := ""
	if deviceName != "" {
		deviceNameMessage = fmt.Sprintf(deviceNameTemplate, deviceName)
	}

	log.Printf(loggingTemplate, device.Address, device.RSSI, deviceNameMessage)
}
