package commands

import (
	"errors"

	"github.com/timmydo/te/buffer"
	"github.com/timmydo/te/widgets"
)

type InsertText struct {
}

func init() {
	register(InsertText{})
}

func (InsertText) Aliases() []string {
	return []string{"insert-text"}
}

func (InsertText) Complete(*widgets.Window, []string) []string {
	return nil
}

func (cmd InsertText) Execute(w *widgets.Window, args []string) error {
	if len(args) < 1 {
		return errors.New("insert-text: Missing arguments")
	}
	w.OpenBuffer.TakeSnapshot(true)
	w.OpenBuffer.Mark = buffer.Loc{-1, -1}
	newPoint := w.OpenBuffer.Data.Contents.InsertString(w.OpenBuffer.Point, args[0])
	w.OpenBuffer.Point = newPoint
	return nil
}
