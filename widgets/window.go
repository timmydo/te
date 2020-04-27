package widgets

import (
	"github.com/timmydo/te/interfaces"
)

type Window struct {
	name          string
	rootDirectory string
	openBuffer    interfaces.Buffer
	clipboard     interfaces.Clipboard
}

func (w Window) Clipboard() interfaces.Clipboard {
	return w.clipboard
}

func (w Window) OpenBuffer() interfaces.Buffer {
	return w.openBuffer
}

func (w Window) SetOpenBuffer(b interfaces.Buffer) {
	w.openBuffer = b
}

func newWindow(name string, rootDirectory string) interfaces.Window {
	w := &Window{name, rootDirectory, interfaces.GetBufferFactory().NewScratchBuffer(), interfaces.GetClipboardProvider().Get()}
	return w
}
