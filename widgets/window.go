package widgets

type Window struct {
	name   string
	Handle interface{}
}

func (app *Window) HandleKeyPress() bool {
	return true
}
