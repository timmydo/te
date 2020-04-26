package commands

import "github.com/timmydo/te/interfaces"

type ScrollUp struct{}
type ScrollDown struct{}

func init() {
	register(ScrollUp{})
	register(ScrollDown{})
}

func (ScrollUp) Aliases() []string {
	return []string{"scroll-page-up"}
}

func (ScrollUp) Complete(interfaces.Window, []string) []string {
	return nil
}

func (cmd ScrollUp) Execute(w interfaces.Window, args []string) error {
	buf := w.OpenBuffer()
	newY := buf.ScrollPosition() - buf.LinesInDisplay()
	if newY < 0 {
		newY = 0
	}
	buf.SetScrollPosition(newY)
	return nil
}

func (ScrollDown) Aliases() []string {
	return []string{"scroll-page-down"}
}

func (ScrollDown) Complete(interfaces.Window, []string) []string {
	return nil
}

func (cmd ScrollDown) Execute(w interfaces.Window, args []string) error {
	buf := w.OpenBuffer()
	newY := buf.ScrollPosition() + buf.LinesInDisplay()
	endY := buf.GetLines().End().Y
	if newY > endY {
		newY = endY - 1
	}
	buf.SetScrollPosition(newY)
	return nil
}
