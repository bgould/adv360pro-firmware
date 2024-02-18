//go:build tinygo && nrf52840

package ble

import (
	"sync"

	"tinygo.org/x/bluetooth"
)

type Mode int

const (
	ModeIdle Mode = iota
	ModePrimary
	ModeSecondary
)

var Default = &BLE{}

type BLE struct {
	adapter *bluetooth.Adapter

	mode     Mode
	modeLock sync.RWMutex

	remoteFound  bool
	remoteDevice bluetooth.ScanResult

	primaryRx bluetooth.DeviceCharacteristic

	secondaryTx bluetooth.Characteristic

	txBuf [20]byte
	rxBuf [20]byte

	SplitName string
}

func (ble *BLE) Enable() error {
	if ble.adapter == nil {
		adapter := bluetooth.DefaultAdapter
		if err := adapter.Enable(); err != nil {
			return err
		}
		if ble.SplitName == "" {
			ble.SplitName = "adv360pro-right"
		}
		ble.adapter = adapter
	}
	return nil
}

func (ble *BLE) Mode() Mode {
	ble.modeLock.RLock()
	defer ble.modeLock.RUnlock()
	return ble.mode
}

func (ble *BLE) idle() {
	println("switching to idle")
	ble.mode = ModeIdle
	ble.primaryRx = bluetooth.DeviceCharacteristic{}
	ble.secondaryTx = bluetooth.Characteristic{}
}

func (ble *BLE) EnablePrimary(notificationCallback func(b []byte)) error {

	ble.modeLock.Lock()
	defer ble.modeLock.Unlock()

	// switch to Idle mode to reset everything if necessary
	ble.idle()

	println("scanning for split")
	if err := ble.adapter.Scan(ble.matchSplit); err != nil {
		return err
	}
	println("connecting to split")
	device, err := ble.adapter.Connect(ble.remoteDevice.Address, bluetooth.ConnectionParams{})
	if err != nil {
		return err
	}
	println("discovering services")
	services, err := device.DiscoverServices(
		[]bluetooth.UUID{bluetooth.ServiceUUIDNordicUART})
	if err != nil {
		return err
	}
	service := services[0]
	println("discovering characteristics")
	chars, err := service.DiscoverCharacteristics(
		[]bluetooth.UUID{bluetooth.CharacteristicUUIDUARTTX})
	if err != nil {
		return err
	}
	ble.primaryRx = chars[0]

	uuid := ble.primaryRx.UUID()
	println("found rx characteristic", uuid.String())

	println("enabling notifications")
	ble.primaryRx.EnableNotifications(notificationCallback)

	println("switching mode to primary")
	ble.mode = ModePrimary
	return nil

}

func (ble *BLE) EnableSecondary() error {

	ble.modeLock.Lock()
	defer ble.modeLock.Unlock()

	// switch to Idle mode to reset everything if necessary
	ble.idle()

	adv := ble.adapter.DefaultAdvertisement()
	err := adv.Configure(bluetooth.AdvertisementOptions{
		LocalName: ble.SplitName,
	})
	err = adv.Start()
	if err != nil {
		return err
	}

	var buf = make([]byte, 0, 3)

	ble.adapter.AddService(&bluetooth.Service{
		UUID: bluetooth.ServiceUUIDNordicUART,
		Characteristics: []bluetooth.CharacteristicConfig{
			{
				Handle: &ble.secondaryTx,
				UUID:   bluetooth.CharacteristicUUIDUARTTX,
				Value:  buf[:],
				Flags:  bluetooth.CharacteristicReadPermission | bluetooth.CharacteristicNotifyPermission,
			},
		},
	})

	println("switching to secondary mode")
	ble.mode = ModeSecondary

	return nil

}

func (ble *BLE) Tx(packet []byte) (ok bool, err error) {
	switch ble.Mode() {
	case ModePrimary:
		return false, nil
	case ModeSecondary:
		if _, err := ble.secondaryTx.Write(packet); err != nil {
			return false, err
		}
		return true, nil
	default:
		return false, nil
	}
}

func (ble *BLE) matchSplit(adapter *bluetooth.Adapter, result bluetooth.ScanResult) {
	splitName := ble.SplitName
	localName := result.LocalName()
	address := result.Address.String()
	println("scanning for", splitName, "-->", localName, address)
	if result.LocalName() != splitName {
		return
	}
	ble.remoteDevice = result
	// Stop the scan.
	err := adapter.StopScan()
	if err != nil {
		println("error stopping scan", err.Error())
	}
	println("found device -->", localName, address)
}

// go func() {
// 	println("starting transmissions")
// 	var b [20]byte
// 	for {
// 		time.Sleep(time.Second)
// 		if ble.Mode() != ModeSecondary {
// 			println("ending transmissions")
// 			return
// 		}
// 		now := time.Now().Unix()
// 		b[0] = byte(now >> 56)
// 		b[1] = byte(now >> 48)
// 		b[2] = byte(now >> 40)
// 		b[3] = byte(now >> 32)
// 		b[4] = byte(now >> 24)
// 		b[5] = byte(now >> 16)
// 		b[6] = byte(now >> 8)
// 		b[7] = byte(now >> 0)
// 		println("send --> ", b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7])
// 		if _, err := ble.secondaryTx.Write(b[0:8]); err != nil {
// 			println("tx error:", err.Error())
// 		}
// 	}
// }()
