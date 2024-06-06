//go:build js

package entrypoint

import (
	"syscall/js"

	"github.com/rabraghib/darijascript/cmd"
	"github.com/rabraghib/darijascript/src/interpreter"
)

func Entrypoint() {
	js.Global().Set("runDarijaScript", js.FuncOf(wasmRunner))

	select {}
}

func wasmRunner(this js.Value, p []js.Value) interface{} {
	if len(p) != 1 {
		return "runDarijaScript() takes exactly 1 argument, " + string(len(p)) + " given"
	}
	sourceCode := p[0].String()
	eval := interpreter.NewEvaluator()
	cmd.RunCode(sourceCode, eval)

	return nil
}
