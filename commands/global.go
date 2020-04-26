package commands

import "github.com/timmydo/te/interfaces"

var (
	GlobalCommands *interfaces.Commands
)

func register(cmd interfaces.Command) {
	if GlobalCommands == nil {
		GlobalCommands = interfaces.NewCommands()
	}
	GlobalCommands.Register(cmd)
}
