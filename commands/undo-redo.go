package commands

import (
	"github.com/timmydo/te/buffer"
	"github.com/timmydo/te/widgets"
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

func (Undo) Complete(*widgets.Window, []string) []string {
	return nil
}

func (cmd Undo) Execute(w *widgets.Window, args []string) error {
	b := w.OpenBuffer
	snap, ok := b.UndoHistory.Value.(*buffer.BufferSnapshot)
	if ok && snap != nil {
		b.UndoHistory = b.UndoHistory.Prev()
		currentState := b.Snapshot(snap.ModifiesBuffer())
		b.RestoreSnapshot(snap)
		b.RedoHistory = b.RedoHistory.Next()
		b.RedoHistory.Value = currentState
	}

	return nil
}

func (Redo) Aliases() []string {
	return []string{"redo"}
}

func (Redo) Complete(*widgets.Window, []string) []string {
	return nil
}

func (cmd Redo) Execute(w *widgets.Window, args []string) error {
	b := w.OpenBuffer
	snap, ok := b.RedoHistory.Value.(*buffer.BufferSnapshot)
	if ok && snap != nil {
		b.RedoHistory = b.RedoHistory.Prev()
		currentState := b.Snapshot(snap.ModifiesBuffer())
		b.RestoreSnapshot(snap)
		b.UndoHistory = b.UndoHistory.Next()
		b.UndoHistory.Value = currentState
	}

	return nil
}
