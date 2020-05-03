package commands

import (
	"github.com/timmydo/te/interfaces"
)

type KillBuffer struct{}

func init() {
	register(KillBuffer{})
}

func (KillBuffer) Aliases() []string {
	return []string{"kill-buffer"}
}

func (KillBuffer) Complete(interfaces.Window, []string) []string {
	return nil
}

func (cmd KillBuffer) Execute(w interfaces.Window, args []string) error {
	// TODO: specify name in args
	buf := w.OpenBuffer()
	bf := interfaces.GetBufferFactory()
	_, current := bf.DeleteBuffer(buf)
	if current == nil {
		current = interfaces.GetBufferFactory().NewScratchBuffer()
	}

	w.SetOpenBuffer(current)
	return nil
}
