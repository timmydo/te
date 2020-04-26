package input

func makeEditMode() *InputCommandSet {
	commands := map[string]*CommandBinding{}
	commands["up"] = makeBinding([]string{"move-point-up-line"})
	commands["Ctrl-p"] = makeBinding([]string{"move-point-up-line"})
	commands["down"] = makeBinding([]string{"move-point-down-line"})
	commands["Ctrl-n"] = makeBinding([]string{"move-point-down-line"})
	commands["return"] = makeBinding([]string{"insert-text", "\n"})
	commands["pageup"] = makeBinding([]string{"scroll-page-up"})
	commands["pagedown"] = makeBinding([]string{"scroll-page-down"})

	addInsertCommands(commands, "insert-text")
	addSingleLineEditCommands(commands)

	return &InputCommandSet{commands}
}

func init() {
	addMode("edit", makeEditMode())
}
