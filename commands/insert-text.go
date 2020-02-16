package commands

import (
	"github.com/timmydo/te/widgets"
	"log"
)


type InsertText struct {
	Text string
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
	log.Printf("cmd insert-text: %v\n", cmd)
	newPoint := w.OpenBuffer.Data.Contents.InsertString(w.OpenBuffer.Point, cmd.Text)
	w.OpenBuffer.Point = newPoint
	return nil
}
