package widgets

import (
	"github.com/timmydo/te/input"
)

type Window struct {
	name string
}

func (app *Window) HandleKeyPress(kp *input.KeyPressInfo) bool {
	return true
}
