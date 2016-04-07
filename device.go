package codova

import (
	"github.com/Archs/js/dom"
	"github.com/gopherjs/gopherjs/js"
)

var (
	defaultDeviceInfo *DeviceInfo
)

type DeviceInfo struct {
	*js.Object
	Cordova      string `js:"cordova"`
	Model        string `js:"model"`
	Platform     string `js:"platform"`
	Uuid         string `js:"uuid"`
	Version      string `js:"version"`
	Manufacturer string `js:"manufacturer"`
	// whether the device is running on a simulator.
	IsVirtual bool `js:"isVirtual"`
	// Get the device hardware serial number
	Serial string `js:"serial"`
}

func Device() *DeviceInfo {
	if defaultDeviceInfo != nil {
		return defaultDeviceInfo
	}
	defaultDeviceInfo = &DeviceInfo{
		Object: js.Global.Get("device"),
	}
	return defaultDeviceInfo
}

// document.addEventListener("deviceready", onDeviceReady, false);
// function onDeviceReady() {
//     console.log(device.cordova);
// }

func OnDeviceReady(fn func()) {
	dom.Document().AddEventListener("deviceready", func(*dom.Event) {
		fn()
	})
}
