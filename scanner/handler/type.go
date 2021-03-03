package handler

import "tinygo.org/x/bluetooth"

type BluetoothHandler interface {
	OnScan(adapter *bluetooth.Adapter, device bluetooth.ScanResult)
}
