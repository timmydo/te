package commands

import (
	"github.com/timmydo/te/buffer"
	"github.com/timmydo/te/widgets"
	"log"
)

type CutText struct{}
type CopyText struct{}
type PasteText struct{}

func init() {
	register(CutText{})
	register(CopyText{})
	register(PasteText{})
}

func getSelection(b *buffer.Buffer) (buffer.Loc, buffer.Loc) {
	if b.Point.LessThan(b.Mark) {
		return b.Point, b.Mark.Move(1, b)
	} else {
		return b.Mark, b.Point
	}
}

func (CutText) Aliases() []string {
	return []string{"cut-text"}
}

func (CutText) Complete(*widgets.Window, []string) []string {
	return nil
}

func (cmd CutText) Execute(w *widgets.Window, args []string) error {
	w.OpenBuffer.TakeSnapshot(true)
	if w.OpenBuffer.Mark.Y == -1 {
		ypos := w.OpenBuffer.Point.Y
		w.Clipboard.SetData(string(w.OpenBuffer.Data.Contents.LineBytes(ypos)) + string(w.OpenBuffer.Data.Contents.Endings))
		w.OpenBuffer.Data.Contents.DeleteLine(ypos)
		log.Printf("point %v\n", w.OpenBuffer.Point)
		log.Printf("end %v\n", w.OpenBuffer.Data.Contents.End())
		w.OpenBuffer.Point = w.OpenBuffer.Point.MoveInBounds(0, w.OpenBuffer)
		w.OpenBuffer.Point.X = 0
		return nil

	}

	start, end := getSelection(w.OpenBuffer)
	data := w.OpenBuffer.Data.Contents.Substr(start, end)
	w.Clipboard.SetData(string(data))
	w.OpenBuffer.Data.Contents.Remove(start, end)
	w.OpenBuffer.Point = start
	w.OpenBuffer.Mark = buffer.Loc{-1, -1}
	return nil
}

func (CopyText) Aliases() []string {
	return []string{"copy-text"}
}

func (CopyText) Complete(*widgets.Window, []string) []string {
	return nil
}

func (cmd CopyText) Execute(w *widgets.Window, args []string) error {
	if w.OpenBuffer.Mark.Y == -1 {
		w.Clipboard.SetData(string(w.OpenBuffer.Data.Contents.LineBytes(w.OpenBuffer.Point.Y)))
		return nil
	}

	w.OpenBuffer.TakeSnapshot(false)
	start, end := getSelection(w.OpenBuffer)
	data := w.OpenBuffer.Data.Contents.Substr(start, end)
	w.Clipboard.SetData(string(data))
	w.OpenBuffer.Mark = buffer.Loc{-1, -1}
	return nil
}

func (PasteText) Aliases() []string {
	return []string{"paste-text"}
}

func (PasteText) Complete(*widgets.Window, []string) []string {
	return nil
}

func (cmd PasteText) Execute(w *widgets.Window, args []string) error {
	// if something is already selected, delete it
	w.OpenBuffer.TakeSnapshot(true)
	if w.OpenBuffer.Mark.Y != -1 {
		start, end := getSelection(w.OpenBuffer)
		w.OpenBuffer.Data.Contents.Remove(start, end)
		w.OpenBuffer.Mark = buffer.Loc{-1, -1}
	}

	newPoint := w.OpenBuffer.Data.Contents.InsertString(w.OpenBuffer.Point, w.Clipboard.GetData())
	w.OpenBuffer.Point = newPoint
	return nil
}
