package widgets

import (
	"log"

	"github.com/timmydo/te/interfaces"
)

func init() {
	interfaces.SetApplication(&application{})
}

type application struct {
	name    string
	windows []interfaces.Window
}

func (app *application) Windows() []interfaces.Window {
	return app.windows
}

func (app *application) CreateWindow(name string, rootDirectory string) interfaces.Window {
	w := NewWindow(name, rootDirectory)
	app.windows = append(app.windows, w)
	log.Printf("Created window %s\n", name)
	return w
}

func (app *application) FindWindow(handle interfaces.Window) interfaces.Window {
	for _, element := range app.windows {
		if element == handle {
			return element
		}
	}

	return nil
}

func remove(s []interface{}, i int) []interface{} {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (app *application) KillWindow(teW interfaces.Window) {
	for i, element := range app.windows {
		if element == teW {
			app.windows[i] = app.windows[len(app.windows)-1]
			app.windows = app.windows[:len(app.windows)-1]
		}
	}
}
