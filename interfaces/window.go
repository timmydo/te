package interfaces

type Window interface {
	OpenBuffer() Buffer
	SetOpenBuffer(Buffer)
	Clipboard() Clipboard
}
