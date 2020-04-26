package commands

import (
	"github.com/timmydo/te/interfaces"
	"github.com/timmydo/te/linearray"
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

func (SetMarkAtPoint) Complete(interfaces.Window, []string) []string {
	return nil
}

func (cmd SetMarkAtPoint) Execute(w interfaces.Window, args []string) error {
	buf := w.OpenBuffer()
	buf.TakeSnapshot(false)
	buf.SetMark(buf.Point())
	return nil
}

func (ClearMark) Aliases() []string {
	return []string{"clear-mark"}
}

func (ClearMark) Complete(interfaces.Window, []string) []string {
	return nil
}

func (cmd ClearMark) Execute(w interfaces.Window, args []string) error {
	buf := w.OpenBuffer()
	buf.SetMark(linearray.Loc{-1, -1})
	return nil
}
