package interfaces

type Application interface {
	Windows() []Window
	CreateWindow(string, string) Window
	FindWindow(Window) Window
	KillWindow(Window)
}


var ap Application

func GetApplication() Application {
	return ap
}

func SetApplication(ap2 Application) {
	ap = ap2
}
