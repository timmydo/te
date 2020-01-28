package widgets

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

func (app *Application) CreateWindow(name string, handle interface{}) *Window {
	w := &Window{name, handle}
	app.Windows = append(app.Windows, w)
	return w
}

func (app *Application) FindWindow(handle interface{}) *Window {
	for _, element := range app.Windows {
		if element.Handle == handle {
			return element
		}
	}

	return nil
}

func remove(s []interface{}, i int) []interface{} {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (app *Application) KillWindow(handle interface{}) {
	for i, element := range app.Windows {
		if element.Handle == handle {
			app.Windows[i] = app.Windows[len(app.Windows)-1]
			app.Windows = app.Windows[:len(app.Windows)-1]
		}
	}
}
