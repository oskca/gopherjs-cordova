// Package callnumber is a GophjerJS wrapper of https://github.com/Rohfosho/CordovaCallNumberPlugin
package callnumber

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	defaultCallNumberHodler *Caller
)

// window.plugins.CallNumber.callNumber(onSuccess, onError, number, bypassAppChooser);

type Caller struct {
	*js.Object
	call func(args ...interface{}) `js:"callNumber"`
}

func Get() *Caller {
	if defaultCallNumberHodler != nil {
		return defaultCallNumberHodler
	}
	defaultCallNumberHodler = &Caller{
		Object: js.Global.Get("plugins").Get("CallNumber"),
	}
	return defaultCallNumberHodler
}

// callbacks: onSuccess, onError
func (c *Caller) Call(tel string, callbacks ...func()) {
	onSuccess := func() {}
	onError := func() {}
	if len(callbacks) >= 1 {
		onSuccess = callbacks[0]
	}
	if len(callbacks) >= 2 {
		onError = callbacks[1]
	}
	c.call(onSuccess, onError, tel, true)
}
