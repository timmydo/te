package interfaces


type Clipboard interface {
	GetData() string
	SetData(string)
}

type ClipboardProvider interface {
	Get() Clipboard
}

var clipboardFactory ClipboardProvider

func GetClipboardProvider() ClipboardProvider {
	return clipboardFactory
}

func SetClipboardProvider(bf ClipboardProvider) {
	clipboardFactory = bf
}
