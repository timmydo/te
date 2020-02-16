package commands

import (
	"github.com/timmydo/te/widgets"
)

type DeleteTextForward struct{}
type DeleteTextBackward struct{}

func init() {
	register(DeleteTextForward{})
	register(DeleteTextBackward{})
}

func (DeleteTextForward) Aliases() []string {
	return []string{"delete-text-forward"}
}

func (DeleteTextForward) Complete(*widgets.Window, []string) []string {
	return nil
}

func (cmd DeleteTextForward) Execute(w *widgets.Window, args []string) error {
	start := w.OpenBuffer.Point.MoveInBounds(0, w.OpenBuffer)
	end := start.MoveInBounds(1, w.OpenBuffer)
	w.OpenBuffer.Data.Contents.Remove(start, end)
	return nil
}

func (DeleteTextBackward) Aliases() []string {
	return []string{"delete-text-backward"}
}

func (DeleteTextBackward) Complete(*widgets.Window, []string) []string {
	return nil
}

func (cmd DeleteTextBackward) Execute(w *widgets.Window, args []string) error {
	end := w.OpenBuffer.Point.MoveInBounds(0, w.OpenBuffer)
	start := end.MoveInBounds(-1, w.OpenBuffer)
	w.OpenBuffer.Data.Contents.Remove(start, end)
	w.OpenBuffer.Point = start
	return nil
}
