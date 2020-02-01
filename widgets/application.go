package widgets

import "log"

var (
	ApplicationInstance Application
)

func init() {
	ApplicationInstance = Application{}
}

type Application struct {
	name    string
	Windows []*Window
}

func (app *Application) CreateWindow(name string) *Window {
	w := NewWindow(name)
	app.Windows = append(app.Windows, w)
	log.Printf("Created window %s\n", name)
	return w
}

func (app *Application) FindWindow(handle *Window) *Window {
	for _, element := range app.Windows {
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

func (app *Application) KillWindow(teW *Window) {
	log.Printf("Destroy window %s\n", teW.name)
	for i, element := range app.Windows {
		if element == teW {
			app.Windows[i] = app.Windows[len(app.Windows)-1]
			app.Windows = app.Windows[:len(app.Windows)-1]
		}
	}
}
