package widgets

import (
	"github.com/timmydo/te/buffer"
	"github.com/timmydo/te/input"
)

type Window struct {
	name                  string
	rootDirectory         string
	OpenBuffer            *buffer.Buffer
	LeftPanelWidthPercent float64
}

func NewWindow(name string, rootDirectory string) *Window {
	w := &Window{name, rootDirectory, buffer.GetScratchBuffer(), float64(30)}
	return w
}

func (app *Window) HandleKeyPress(kp *input.KeyPressInfo) bool {
	return true
}
