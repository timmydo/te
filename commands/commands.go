package commands

import (
	"errors"
	"sort"
	"strings"

	"github.com/google/shlex"

	"github.com/timmydo/te/widgets"
)

type Command interface {
	Aliases() []string
	Execute(*widgets.Window, []string) error
	Complete(*widgets.Window, []string) []string
}

type Commands map[string]Command

func NewCommands() *Commands {
	cmds := Commands(make(map[string]Command))
	return &cmds
}

func (cmds *Commands) dict() map[string]Command {
	return map[string]Command(*cmds)
}

func (cmds *Commands) Names() []string {
	names := make([]string, 0)

	for k := range cmds.dict() {
		names = append(names, k)
	}
	return names
}

func (cmds *Commands) Register(cmd Command) {
	if len(cmd.Aliases()) < 1 {
		return
	}
	for _, alias := range cmd.Aliases() {
		cmds.dict()[alias] = cmd
	}
}

type NoSuchCommand string

func (err NoSuchCommand) Error() string {
	return "Unknown command " + string(err)
}

func (cmds *Commands) ExecuteCommand(win *widgets.Window, args []string) error {
	if len(args) == 0 {
		return errors.New("Expected a command.")
	}
	if cmd, ok := cmds.dict()[args[0]]; ok {
		return cmd.Execute(win, args)
	}
	return NoSuchCommand(args[0])
}

func (cmds *Commands) GetCompletions(win *widgets.Window, cmd string) []string {
	args, err := shlex.Split(cmd)
	if err != nil {
		return nil
	}

	if len(args) == 0 {
		names := cmds.Names()
		sort.Strings(names)
		return names
	}

	if len(args) > 1 || cmd[len(cmd)-1] == ' ' {
		if cmd, ok := cmds.dict()[args[0]]; ok {
			var completions []string
			if len(args) > 1 {
				completions = cmd.Complete(win, args[1:])
			} else {
				completions = cmd.Complete(win, []string{})
			}
			if completions != nil && len(completions) == 0 {
				return nil
			}

			options := make([]string, 0)
			for _, option := range completions {
				options = append(options, args[0]+" "+option)
			}
			return options
		}
		return nil
	}

	names := cmds.Names()
	options := make([]string, 0)
	for _, name := range names {
		if strings.HasPrefix(name, args[0]) {
			options = append(options, name)
		}
	}

	if len(options) > 0 {
		return options
	}
	return nil
}
