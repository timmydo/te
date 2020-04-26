package widgets

import (
	"github.com/timmydo/te/buffer"
)

type Window struct {
	name          string
	rootDirectory string
	OpenBuffer    *buffer.Buffer
	Clipboard     *buffer.Clipboard
}

func NewWindow(name string, rootDirectory string) *Window {
	w := &Window{name, rootDirectory, buffer.GetScratchBuffer(), buffer.GetClipboard()}
	return w
}
