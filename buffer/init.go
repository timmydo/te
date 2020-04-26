package buffer

import (
	"github.com/timmydo/te/interfaces"
)

func Initialize() {
	interfaces.SetBufferFactory(&myBufferFactory{})
	interfaces.SetClipboardProvider(&clipboardProvider{})
}
