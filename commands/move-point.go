package commands

import (
	"log"

	"github.com/timmydo/te/interfaces"
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

func (MovePointLeftChar) Complete(interfaces.Window, []string) []string {
	return nil
}

func (MovePointLeftChar) Execute(w interfaces.Window, args []string) error {
	buf := w.OpenBuffer()
	// if past end of line, move it to the line
	buf.SetPoint(buf.Point().MoveInBounds(0, buf.GetLines()).MoveInBounds(-1, buf.GetLines()))
	return nil
}

func (MovePointRightChar) Aliases() []string {
	return []string{"move-point-right-char"}
}

func (MovePointRightChar) Complete(interfaces.Window, []string) []string {
	return nil
}

func (MovePointRightChar) Execute(w interfaces.Window, args []string) error {
	buf := w.OpenBuffer()
	buf.SetPoint(buf.Point().MoveInBounds(0, buf.GetLines()).MoveInBounds(1, buf.GetLines()))
	log.Printf("After r-char: %v\n", buf.Point)
	return nil
}

func (MovePointUpLine) Aliases() []string {
	return []string{"move-point-up-line"}
}

func (MovePointUpLine) Complete(interfaces.Window, []string) []string {
	return nil
}

func (MovePointUpLine) Execute(w interfaces.Window, args []string) error {
	buf := w.OpenBuffer()
	buf.SetPoint(buf.Point().MoveDownLines(-1, buf.GetLines()))
	log.Printf("After r-char: %v\n", buf.Point)
	return nil
}

func (MovePointDownLine) Aliases() []string {
	return []string{"move-point-down-line"}
}

func (MovePointDownLine) Complete(interfaces.Window, []string) []string {
	return nil
}

func (MovePointDownLine) Execute(w interfaces.Window, args []string) error {
	buf := w.OpenBuffer()
	buf.SetPoint(buf.Point().MoveDownLines(1, buf.GetLines()))
	log.Printf("After r-char: %v\n", buf.Point)
	return nil
}

func (MovePointStartOfLine) Aliases() []string {
	return []string{"move-point-start-of-line"}
}

func (MovePointStartOfLine) Complete(interfaces.Window, []string) []string {
	return nil
}

func (MovePointStartOfLine) Execute(w interfaces.Window, args []string) error {
	buf := w.OpenBuffer()
	buf.SetPoint(buf.Point().MoveStartOfLine(buf.GetLines()))
	log.Printf("After r-char: %v\n", buf.Point)
	return nil
}

func (MovePointEndOfLine) Aliases() []string {
	return []string{"move-point-end-of-line"}
}

func (MovePointEndOfLine) Complete(interfaces.Window, []string) []string {
	return nil
}

func (MovePointEndOfLine) Execute(w interfaces.Window, args []string) error {
	buf := w.OpenBuffer()
	buf.SetPoint(buf.Point().MoveEndOfLine(buf.GetLines()))
	log.Printf("After r-char: %v\n", buf.Point)
	return nil
}
