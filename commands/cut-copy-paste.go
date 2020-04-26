package commands

import (
	"log"

	"github.com/timmydo/te/interfaces"
	"github.com/timmydo/te/linearray"
)

type CutText struct{}
type CopyText struct{}
type PasteText struct{}

func init() {
	register(CutText{})
	register(CopyText{})
	register(PasteText{})
}

func getSelection(b interfaces.Buffer) (linearray.Loc, linearray.Loc) {
	if b.Point().LessThan(b.Mark()) {
		return b.Point(), b.Mark().Move(1, b.GetLines())
	} else {
		return b.Mark(), b.Point()
	}
}

func (CutText) Aliases() []string {
	return []string{"cut-text"}
}

func (CutText) Complete(interfaces.Window, []string) []string {
	return nil
}

func (cmd CutText) Execute(w interfaces.Window, args []string) error {
	buf := w.OpenBuffer()
	buf.TakeSnapshot(true)
	if buf.Mark().Y == -1 {
		ypos := buf.Point().Y
		w.Clipboard().SetData(string(buf.GetLines().LineBytes(ypos)) + string(buf.GetLines().Endings))
		buf.GetLines().DeleteLine(ypos)
		log.Printf("point %v\n", buf.Point())
		log.Printf("end %v\n", buf.GetLines().End())
		newPt := buf.Point().MoveInBounds(0, buf.GetLines())
		buf.SetPoint(linearray.Loc{0, newPt.Y})
		return nil

	}

	start, end := getSelection(buf)
	data := buf.GetLines().Substr(start, end)
	w.Clipboard().SetData(string(data))
	buf.GetLines().Remove(start, end)
	buf.SetPoint(start)
	buf.SetMark(linearray.Loc{-1, -1})
	return nil
}

func (CopyText) Aliases() []string {
	return []string{"copy-text"}
}

func (CopyText) Complete(interfaces.Window, []string) []string {
	return nil
}

func (cmd CopyText) Execute(w interfaces.Window, args []string) error {
	buf := w.OpenBuffer()
	if buf.Mark().Y == -1 {
		w.Clipboard().SetData(string(buf.GetLines().LineBytes(buf.Point().Y)))
		return nil
	}

	buf.TakeSnapshot(false)
	start, end := getSelection(buf)
	data := buf.GetLines().Substr(start, end)
	w.Clipboard().SetData(string(data))
	buf.SetMark(linearray.Loc{-1, -1})
	return nil
}

func (PasteText) Aliases() []string {
	return []string{"paste-text"}
}

func (PasteText) Complete(interfaces.Window, []string) []string {
	return nil
}

func (cmd PasteText) Execute(w interfaces.Window, args []string) error {
	buf := w.OpenBuffer()
	// if something is already selected, delete it
	buf.TakeSnapshot(true)
	if buf.Mark().Y != -1 {
		start, end := getSelection(buf)
		buf.GetLines().Remove(start, end)
		buf.SetMark(linearray.Loc{-1, -1})
	}

	newPoint := buf.GetLines().InsertString(buf.Point(), w.Clipboard().GetData())
	buf.SetPoint(newPoint)
	return nil
}
