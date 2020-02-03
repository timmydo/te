package commands

import (
	"github.com/timmydo/te/widgets"
	"log"
)

type MovePointLeftChar struct{}
type MovePointRightChar struct{}

func init() {
	register(MovePointLeftChar{})
	register(MovePointRightChar{})
}

func (MovePointLeftChar) Aliases() []string {
	return []string{"move-point-left-char"}
}

func (MovePointLeftChar) Complete(*widgets.Window, []string) []string {
	return nil
}

func (MovePointLeftChar) Execute(w *widgets.Window, args []string) error {
	log.Printf("Before left-char: %v\n", w.OpenBuffer.Point)
	w.OpenBuffer.Point = w.OpenBuffer.Point.Move(-1, w.OpenBuffer)
	log.Printf("After left-char: %v\n", w.OpenBuffer.Point)
	return nil
}

func (MovePointRightChar) Aliases() []string {
	return []string{"move-point-right-char"}
}

func (MovePointRightChar) Complete(*widgets.Window, []string) []string {
	return nil
}

func (MovePointRightChar) Execute(w *widgets.Window, args []string) error {
	log.Printf("Before r-char: %v\n", w.OpenBuffer.Point)
	w.OpenBuffer.Point = w.OpenBuffer.Point.Move(1, w.OpenBuffer)
	log.Printf("After r-char: %v\n", w.OpenBuffer.Point)
	return nil
}
