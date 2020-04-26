package interfaces

type Window interface {
	OpenBuffer() Buffer
	Clipboard() Clipboard
}
