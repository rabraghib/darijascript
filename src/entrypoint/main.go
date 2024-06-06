//go:build !js

package entrypoint

import "github.com/rabraghib/darijascript/cmd"

func Entrypoint() {
	cmd.Execute()
}
