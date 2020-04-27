package widgets

import "github.com/timmydo/te/interfaces"

func Initialize() {
	interfaces.SetApplication(&application{})
}
