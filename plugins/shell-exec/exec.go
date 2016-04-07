// // Package shell is a GophjerJS wrapper of https://github.com/petervojtek/cordova-plugin-shell-exec
package shell

import (
	"fmt"
	"github.com/gopherjs/gopherjs/js"
)

type ExecResult struct {
	*js.Object
	ExitStatus int    `js:"exitStatus"`
	Output     string `js:"output"`
}

func ExecRaw(cmd string, cb func(res *ExecResult)) {
	js.Global.Get("ShellExec").Call("exec", cmd, cb)
}

func ExecOutput(cmd string) (output string, err error) {
	ch := make(chan *ExecResult, 0)
	ExecRaw(cmd, func(res *ExecResult) {
		ch <- res
	})
	res := <-ch
	if res.ExitStatus == 0 {
		return res.Output, nil
	}
	return res.Output, fmt.Errorf("status:%d, msg:%s", res.ExitStatus, res.Output)
}

func Exec(cmd string) error {
	_, err := ExecOutput(cmd)
	return err
}
