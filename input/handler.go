package input

import "log"

type InputCommandSet struct {
	Commands map[string]*CommandBinding
}

var (
	modeMap map[string]*InputCommandSet
)

type CommandBinding struct {
	Args []string
}

func makeBinding(args []string) *CommandBinding {
	return &CommandBinding{args}
}

func makeEditMode() *InputCommandSet {
	commands := map[string]*CommandBinding{}
	commands["left"] = makeBinding([]string{"move-point-left-char"})
	commands["Ctrl-b"] = makeBinding([]string{"move-point-left-char"})
	commands["right"] = makeBinding([]string{"move-point-right-char"})
	commands["Ctrl-f"] = makeBinding([]string{"move-point-right-char"})
	commands["up"] = makeBinding([]string{"move-point-up-line"})
	commands["Ctrl-p"] = makeBinding([]string{"move-point-up-line"})
	commands["down"] = makeBinding([]string{"move-point-down-line"})
	commands["Ctrl-n"] = makeBinding([]string{"move-point-down-line"})
	commands["Ctrl-d"] = makeBinding([]string{"delete-text-forward"})
	commands["delete"] = makeBinding([]string{"delete-text-forward"})
	commands["backspace"] = makeBinding([]string{"delete-text-backward"})
	commands["return"] = makeBinding([]string{"insert-text", "\n"})
	commands["space"] = makeBinding([]string{"insert-text", " "})

	commands["home"] = makeBinding([]string{"move-point-start-of-line"})
	commands["end"] = makeBinding([]string{"move-point-end-of-line"})

	commands["Ctrl-a"] = makeBinding([]string{"move-point-start-of-line"})
	commands["Ctrl-e"] = makeBinding([]string{"move-point-end-of-line"})

	commands["Ctrl-x"] = makeBinding([]string{"cut-text"})
	commands["Ctrl-c"] = makeBinding([]string{"copy-text"})
	commands["Ctrl-v"] = makeBinding([]string{"paste-text"})

	commands["Ctrl-z"] = makeBinding([]string{"undo"})
	commands["Ctrl-y"] = makeBinding([]string{"redo"})

	commands["Ctrl-g"] = makeBinding([]string{"clear-mark"})
	commands["Ctrl-space"] = makeBinding([]string{"set-mark-at-point"})

	commands["pageup"] = makeBinding([]string{"scroll-page-up"})
	commands["pagedown"] = makeBinding([]string{"scroll-page-down"})

	for _, key := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()=+-_/?\\|\"'<>,.~`[]{}" {
		commands[string(key)] = makeBinding([]string{"insert-text", string(key)})
	}

	return &InputCommandSet{commands}
}

func init() {
	editMode := makeEditMode()
	modeMap = map[string]*InputCommandSet{}
	modeMap["edit"] = editMode
}

func FindCommand(kp *KeyPressInfo, mode string) []string {
	kpName := kp.GetName()
	log.Printf("Mode %s, Key press: %s\n", mode, kpName)

	if mode, modeFound := modeMap[mode]; modeFound {
		if item, found := mode.Commands[kpName]; found {
			return item.Args
		}
	}

	return nil
}
