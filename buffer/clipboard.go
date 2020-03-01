package buffer

import "log"

var (
	GlobalClipboard *Clipboard
)

type Clipboard struct {
	data string
}

func init() {
	GlobalClipboard = &Clipboard{""}
}

func GetClipboard() *Clipboard {
	return GlobalClipboard
}

func (c *Clipboard) SetData(str string) {
	log.Printf("Clipboard: \"%v\"\n", str)
	c.data = str
}

func (c *Clipboard) GetData() string {
	return c.data
}
