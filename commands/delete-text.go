package commands

import "github.com/timmydo/te/interfaces"

type DeleteTextForward struct{}
type DeleteTextBackward struct{}

func init() {
	register(DeleteTextForward{})
	register(DeleteTextBackward{})
}

func (DeleteTextForward) Aliases() []string {
	return []string{"delete-text-forward"}
}

func (DeleteTextForward) Complete(interfaces.Window, []string) []string {
	return nil
}

func (cmd DeleteTextForward) Execute(w interfaces.Window, args []string) error {
	buf := w.OpenBuffer()
	buf.TakeSnapshot(true)
	start := buf.Point().MoveInBounds(0, buf.GetLines())
	end := start.MoveInBounds(1, buf.GetLines())
	buf.GetLines().Remove(start, end)
	return nil
}

func (DeleteTextBackward) Aliases() []string {
	return []string{"delete-text-backward"}
}

func (DeleteTextBackward) Complete(interfaces.Window, []string) []string {
	return nil
}

func (cmd DeleteTextBackward) Execute(w interfaces.Window, args []string) error {
	buf := w.OpenBuffer()
	buf.TakeSnapshot(true)
	end := buf.Point().MoveInBounds(0, buf.GetLines())
	start := end.MoveInBounds(-1, buf.GetLines())
	buf.GetLines().Remove(start, end)
	buf.SetPoint(start)
	return nil
}
