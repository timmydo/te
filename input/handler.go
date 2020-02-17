package input

import "log"

var (
	commands map[string]*CommandBinding
)

type CommandBinding struct {
	Args []string
}

func makeBinding(args []string) *CommandBinding {
	return &CommandBinding{args}
}

func init() {
	commands = map[string]*CommandBinding{}
	commands["left"] = makeBinding([]string{"move-point-left-char"})
	commands["Ctrl-b"] = makeBinding([]string{"move-point-left-char"})
	commands["right"] = makeBinding([]string{"move-point-right-char"})
	commands["Ctrl-f"] = makeBinding([]string{"move-point-right-char"})
	commands["up"] = makeBinding([]string{"move-point-up-line"})
	commands["Ctrl-p"] = makeBinding([]string{"move-point-up-line"})
	commands["down"] = makeBinding([]string{"move-point-down-line"})
	commands["Ctrl-n"] = makeBinding([]string{"move-point-down-line"})
	commands["delete"] = makeBinding([]string{"delete-text-forward"})
	commands["backspace"] = makeBinding([]string{"delete-text-backward"})
	commands["return"] = makeBinding([]string{"insert-text", "\n"})
	commands["space"] = makeBinding([]string{"insert-text", " "})

	commands["home"] = makeBinding([]string{"move-point-start-of-line"})
	commands["end"] = makeBinding([]string{"move-point-end-of-line"})

	commands["Ctrl-a"] = makeBinding([]string{"move-point-start-of-line"})
	commands["Ctrl-e"] = makeBinding([]string{"move-point-end-of-line"})

	commands["Ctrl-g"] = makeBinding([]string{"clear-mark"})
	commands["Ctrl-space"] = makeBinding([]string{"set-mark-at-point"})

	for _, key := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()=+-_/?\\|\"'<>,.~`[]{}" {
		commands[string(key)] = makeBinding([]string{"insert-text", string(key)})
	}
}

func FindCommand(kp *KeyPressInfo) []string {
	kpName := kp.GetName()
	log.Printf("Key press: %s\n", kpName)
	if item, found := commands[kpName]; found {
		return item.Args
	}
	return nil
}
