package commands

import (
	"github.com/timmydo/te/widgets"
)

type ScrollUp struct{}
type ScrollDown struct{}

func init() {
	register(ScrollUp{})
	register(ScrollDown{})
}

func (ScrollUp) Aliases() []string {
	return []string{"scroll-page-up"}
}

func (ScrollUp) Complete(*widgets.Window, []string) []string {
	return nil
}

func (cmd ScrollUp) Execute(w *widgets.Window, args []string) error {
	newY := w.OpenBuffer.ScrollPosition.Y - w.OpenBuffer.LinesInDisplay
	if newY < 0 {
		newY = 0
	}
	w.OpenBuffer.ScrollPosition.Y = newY
	return nil
}

func (ScrollDown) Aliases() []string {
	return []string{"scroll-page-down"}
}

func (ScrollDown) Complete(*widgets.Window, []string) []string {
	return nil
}

func (cmd ScrollDown) Execute(w *widgets.Window, args []string) error {
	newY := w.OpenBuffer.ScrollPosition.Y + w.OpenBuffer.LinesInDisplay
	endY := w.OpenBuffer.Data.Contents.End().Y
	if newY > endY {
		newY = endY - 1
	}
	w.OpenBuffer.ScrollPosition.Y = newY
	return nil
}
