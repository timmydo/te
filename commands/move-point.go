package commands

import (
	"github.com/timmydo/te/widgets"
	"log"
)

type MovePointLeftChar struct{}
type MovePointRightChar struct{}
type MovePointUpLine struct{}
type MovePointDownLine struct{}
type MovePointStartOfLine struct{}
type MovePointEndOfLine struct{}

func init() {
	register(MovePointLeftChar{})
	register(MovePointRightChar{})
	register(MovePointUpLine{})
	register(MovePointDownLine{})
	register(MovePointStartOfLine{})
	register(MovePointEndOfLine{})
}

func (MovePointLeftChar) Aliases() []string {
	return []string{"move-point-left-char"}
}

func (MovePointLeftChar) Complete(*widgets.Window, []string) []string {
	return nil
}

func (MovePointLeftChar) Execute(w *widgets.Window, args []string) error {
	// if past end of line, move it to the line
	w.OpenBuffer.Point = w.OpenBuffer.Point.MoveInBounds(0, w.OpenBuffer).MoveInBounds(-1, w.OpenBuffer)
	return nil
}

func (MovePointRightChar) Aliases() []string {
	return []string{"move-point-right-char"}
}

func (MovePointRightChar) Complete(*widgets.Window, []string) []string {
	return nil
}

func (MovePointRightChar) Execute(w *widgets.Window, args []string) error {
	w.OpenBuffer.Point = w.OpenBuffer.Point.MoveInBounds(0, w.OpenBuffer).MoveInBounds(1, w.OpenBuffer)
	log.Printf("After r-char: %v\n", w.OpenBuffer.Point)
	return nil
}

func (MovePointUpLine) Aliases() []string {
	return []string{"move-point-up-line"}
}

func (MovePointUpLine) Complete(*widgets.Window, []string) []string {
	return nil
}

func (MovePointUpLine) Execute(w *widgets.Window, args []string) error {
	w.OpenBuffer.Point = w.OpenBuffer.Point.MoveDownLines(-1, w.OpenBuffer)
	log.Printf("After r-char: %v\n", w.OpenBuffer.Point)
	return nil
}

func (MovePointDownLine) Aliases() []string {
	return []string{"move-point-down-line"}
}

func (MovePointDownLine) Complete(*widgets.Window, []string) []string {
	return nil
}

func (MovePointDownLine) Execute(w *widgets.Window, args []string) error {
	w.OpenBuffer.Point = w.OpenBuffer.Point.MoveDownLines(1, w.OpenBuffer)
	log.Printf("After r-char: %v\n", w.OpenBuffer.Point)
	return nil
}

func (MovePointStartOfLine) Aliases() []string {
	return []string{"move-point-start-of-line"}
}

func (MovePointStartOfLine) Complete(*widgets.Window, []string) []string {
	return nil
}

func (MovePointStartOfLine) Execute(w *widgets.Window, args []string) error {
	w.OpenBuffer.Point = w.OpenBuffer.Point.MoveStartOfLine(w.OpenBuffer)
	log.Printf("After r-char: %v\n", w.OpenBuffer.Point)
	return nil
}

func (MovePointEndOfLine) Aliases() []string {
	return []string{"move-point-end-of-line"}
}

func (MovePointEndOfLine) Complete(*widgets.Window, []string) []string {
	return nil
}

func (MovePointEndOfLine) Execute(w *widgets.Window, args []string) error {
	w.OpenBuffer.Point = w.OpenBuffer.Point.MoveEndOfLine(w.OpenBuffer)
	log.Printf("After r-char: %v\n", w.OpenBuffer.Point)
	return nil
}
