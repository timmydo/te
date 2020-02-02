package widgets

import (
	"github.com/timmydo/te/buffer"
	"github.com/timmydo/te/input"
)

type Window struct {
	name          string
	rootDirectory string
	OpenBuffer    *buffer.Buffer
}

func NewWindow(name string, rootDirectory string) *Window {
	w := &Window{name, rootDirectory, buffer.GetScratchBuffer()}
	return w
}

func (app *Window) HandleKeyPress(kp *input.KeyPressInfo) bool {
	return true
}
