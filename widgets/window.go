package widgets

import (
	"github.com/timmydo/te/buffer"
	"github.com/timmydo/te/input"
)

type Window struct {
	name       string
	OpenBuffer *buffer.Buffer
}

func NewWindow(name string) *Window {
	w := &Window{name, buffer.GetScratchBuffer()}
	return w
}

func (app *Window) HandleKeyPress(kp *input.KeyPressInfo) bool {
	return true
}
