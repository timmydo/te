package input

type CommandBinding struct {
	Args []string
}

func MakeBinding(args ...string) *CommandBinding {
	return &CommandBinding{args}
}

func AddInsertCommands(commands map[string]*CommandBinding, funcName string) {
	for _, key := range "abcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()=+-_/?\\|\"'<>,.~`[]{}" {
		commands[string(key)] = MakeBinding(funcName, string(key))
	}

	for _, key := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		commands["Shift-"+string(key)] = MakeBinding(funcName, string(key))
	}
}

func AddBufferCommands(commands map[string]*CommandBinding) {
	commands["Ctrl-Shift-k"] = MakeBinding("kill-buffer")
}

func AddSingleLineEditCommands(commands map[string]*CommandBinding) {
	commands["left"] = MakeBinding("move-point-left-char")
	commands["Ctrl-b"] = MakeBinding("move-point-left-char")
	commands["right"] = MakeBinding("move-point-right-char")
	commands["Ctrl-f"] = MakeBinding("move-point-right-char")
	commands["Ctrl-d"] = MakeBinding("delete-text-forward")
	commands["delete"] = MakeBinding("delete-text-forward")
	commands["backspace"] = MakeBinding("delete-text-backward")
	commands["space"] = MakeBinding("insert-text", " ")
	commands["home"] = MakeBinding("move-point-start-of-line")
	commands["end"] = MakeBinding("move-point-end-of-line")
	commands["Ctrl-a"] = MakeBinding("move-point-start-of-line")
	commands["Ctrl-e"] = MakeBinding("move-point-end-of-line")
	commands["Ctrl-x"] = MakeBinding("cut-text")
	commands["Ctrl-c"] = MakeBinding("copy-text")
	commands["Ctrl-v"] = MakeBinding("paste-text")
	commands["Ctrl-z"] = MakeBinding("undo")
	commands["Ctrl-y"] = MakeBinding("redo")
	commands["Ctrl-g"] = MakeBinding("clear-mark")
	commands["Ctrl-space"] = MakeBinding("set-mark-at-point")

}
