package commands

import (
	"github.com/timmydo/te/interfaces"
)

type Undo struct{}
type Redo struct{}

func init() {
	register(Undo{})
	register(Redo{})
}

func (Undo) Aliases() []string {
	return []string{"undo"}
}

func (Undo) Complete(interfaces.Window, []string) []string {
	return nil
}

func (cmd Undo) Execute(w interfaces.Window, args []string) error {
	w.OpenBuffer().Undo()
	return nil
}

func (Redo) Aliases() []string {
	return []string{"redo"}
}

func (Redo) Complete(interfaces.Window, []string) []string {
	return nil
}

func (cmd Redo) Execute(w interfaces.Window, args []string) error {
	w.OpenBuffer().Redo()
	return nil
}
