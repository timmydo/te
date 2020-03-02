package commands

import (
	"github.com/timmydo/te/buffer"
	"github.com/timmydo/te/widgets"
)

type SetMarkAtPoint struct{}
type ClearMark struct{}

func init() {
	register(SetMarkAtPoint{})
	register(ClearMark{})
}

func (SetMarkAtPoint) Aliases() []string {
	return []string{"set-mark-at-point"}
}

func (SetMarkAtPoint) Complete(*widgets.Window, []string) []string {
	return nil
}

func (cmd SetMarkAtPoint) Execute(w *widgets.Window, args []string) error {
	w.OpenBuffer.TakeSnapshot(false)
	w.OpenBuffer.Mark = w.OpenBuffer.Point
	return nil
}

func (ClearMark) Aliases() []string {
	return []string{"clear-mark"}
}

func (ClearMark) Complete(*widgets.Window, []string) []string {
	return nil
}

func (cmd ClearMark) Execute(w *widgets.Window, args []string) error {
	w.OpenBuffer.Mark = buffer.Loc{-1, -1}
	return nil
}
