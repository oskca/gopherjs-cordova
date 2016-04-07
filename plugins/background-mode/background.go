// Package background is a GopherJS wrapper of Cordova Plugin:
// https://github.com/katzer/cordova-plugin-background-mode
package background

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	defaultBackgroundMode *BackgroundMode
)

type BackgroundMode struct {
	*js.Object
	// To activate the background mode the app needs to be in foreground.
	Enable func() `js:"enable"`
	// Once the background mode has been disabled, the app will be paused when in background.
	Disable     func()                 `js:"disable"`
	IsEnabled   func()                 `js:"isEnabled"`
	IsActive    func()                 `js:"isActive"`
	GetDefaults func() *js.Object      `js:"getDefaults"`
	SetDefaults func(conf interface{}) `js:"setDefaults"`
	Configure   func(conf interface{}) `js:"configure"`
	// The backgroundMode.onactivate interface can be used to
	// get notified when the background mode has been activated.
	Onactivate func() `js:"onactivate"`
	// Once the mode has been deactivated the app will be paused
	// soon after the callback has been fired.
	Ondeactivate func() `js:"ondeactivate"`
	// Get informed when the background mode could not been activated
	Onfailure func(errorCode int) `js:"onfailure"`
}

func Get() *BackgroundMode {
	if defaultBackgroundMode != nil {
		return defaultBackgroundMode
	}
	// cordova.plugins.backgroundMode
	o := js.Global.Get("cordova").Get("plugins").Get("backgroundMode")
	defaultBackgroundMode = &BackgroundMode{
		Object: o,
	}
	return defaultBackgroundMode
}

func (b *BackgroundMode) Notify(title, msg string) {
	b.Configure(js.M{
		"title": title,
		"text":  msg,
	})
}
