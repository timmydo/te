package commands

import (
	"errors"

	"github.com/timmydo/te/interfaces"
)

type NewBuffer struct{}

func init() {
	register(NewBuffer{})
}

func (NewBuffer) Aliases() []string {
	return []string{"new-buffer"}
}

func (NewBuffer) Complete(interfaces.Window, []string) []string {
	return nil
}

func (cmd NewBuffer) Execute(w interfaces.Window, args []string) error {
	if len(args) < 1 {
		return errors.New("Missing arg: buffer mode")
	}

	modeName := args[0]
	b := interfaces.GetBufferFactory().CreateBuffer(modeName)
	w.SetOpenBuffer(b)
	return nil
}
