package buffer

import (
	"log"

	"github.com/timmydo/te/interfaces"
)

var (
	globalClipboard *clipboard
)

func init() {
	globalClipboard = &clipboard{""}
}

type clipboard struct {
	data string
}

type clipboardProvider struct {
}

func (c clipboardProvider) Get() interfaces.Clipboard {
	return globalClipboard
}

func (c *clipboard) SetData(str string) {
	log.Printf("Clipboard: \"%v\"\n", str)
	c.data = str
}

func (c *clipboard) GetData() string {
	return c.data
}
