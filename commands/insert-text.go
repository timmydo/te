package commands

import (
	"errors"

	"github.com/timmydo/te/interfaces"
	"github.com/timmydo/te/linearray"
)

type InsertText struct {
}

func init() {
	register(InsertText{})
}

func (InsertText) Aliases() []string {
	return []string{"insert-text"}
}

func (InsertText) Complete(interfaces.Window, []string) []string {
	return nil
}

func (cmd InsertText) Execute(w interfaces.Window, args []string) error {
	buf := w.OpenBuffer()
	if len(args) < 1 {
		return errors.New("insert-text: Missing arguments")
	}
	buf.TakeSnapshot(true)
	buf.SetMark(linearray.Loc{-1, -1})
	newPoint := buf.GetLines().InsertString(buf.Point(), args[0])
	buf.SetPoint(newPoint)
	return nil
}
