package widgets

import (
	"github.com/timmydo/te/input"
)

type Window struct {
	name   string
	Handle interface{}
}

func (app *Window) HandleKeyPress(kp *input.KeyPressInfo) bool {
	return true
}
