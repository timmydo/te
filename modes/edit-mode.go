package modes

import (
	"github.com/timmydo/te/commands"
	"github.com/timmydo/te/input"
	"github.com/timmydo/te/interfaces"
	"github.com/timmydo/te/theme"
)

var (
	bindings map[string]*input.CommandBinding
)

func setupBindings() {
	bindings = map[string]*input.CommandBinding{}
	bindings["up"] = input.MakeBinding("move-point-up-line")
	bindings["Ctrl-p"] = input.MakeBinding("move-point-up-line")
	bindings["down"] = input.MakeBinding("move-point-down-line")
	bindings["Ctrl-n"] = input.MakeBinding("move-point-down-line")
	bindings["return"] = input.MakeBinding("insert-text", "\n")
	bindings["pageup"] = input.MakeBinding("scroll-page-up")
	bindings["pagedown"] = input.MakeBinding("scroll-page-down")

	input.AddInsertCommands(bindings, "insert-text")
	input.AddSingleLineEditCommands(bindings)

}

type editModeFactory struct{}
type editMode struct{}

func (f editModeFactory) Create() interfaces.EditorMode {
	return &editMode{}
}

func (m editMode) Name() string {
	return "edit"
}

func (m editMode) ExecuteCommand(w interfaces.Window, key *input.KeyPressInfo) error {
	if b, found := bindings[key.GetName()]; found {
		return commands.GlobalCommands.ExecuteCommand(w, b.Args)
	}

	return nil
}

func (this *editMode) GetBufferStyle() *theme.BufferThemeStyle {
	return theme.DefaultBufferTheme
}

func (this *editMode) GetLineStyle(int) *theme.LineThemeStyle {
	return theme.DefaultLineTheme
}

func (this *editMode) GetCharacterStyle(int, int) *theme.CharacterThemeStyle {
	return theme.DefaultCharacterTheme
}

func init() {
	setupBindings()
	interfaces.AddMode("edit", &editModeFactory{})
}
