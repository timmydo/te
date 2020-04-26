package modes

import (
	"github.com/timmydo/te/commands"
	"github.com/timmydo/te/input"
	"github.com/timmydo/te/interfaces"
	"github.com/timmydo/te/theme"
)

type findFileModeFactory struct {
	bindings map[string]*input.CommandBinding
}

func (f findFileModeFactory) Create() interfaces.EditorMode {
	return &findFileMode{f.bindings, -1}
}

type findFileMode struct {
	bindings     map[string]*input.CommandBinding
	selectedLine int
}

func (m findFileMode) Name() string {
	return "findfile"
}

func (m findFileMode) ExecuteCommand(w interfaces.Window, key string) error {
	if b, found := m.bindings[key]; found {
		return commands.GlobalCommands.ExecuteCommand(w, b.Args)
	}
	return nil
}

func (this *findFileMode) GetBufferStyle() *theme.BufferThemeStyle {
	return theme.DefaultBufferTheme
}

var (
	UnselectedLineTheme = &theme.LineThemeStyle{
		Background: theme.Color{0.95, 0.95, 0.95, 1.0},
	}
	SelectedLineTheme = &theme.LineThemeStyle{
		Background: theme.Color{0.85, 0.85, 0.85, 1.0},
	}
)

func (this *findFileMode) GetLineStyle(line int) *theme.LineThemeStyle {
	if this.selectedLine == line {
		return SelectedLineTheme
	} else {
		return UnselectedLineTheme
	}
}

func (this *findFileMode) GetCharacterStyle(int, int) *theme.CharacterThemeStyle {
	return theme.DefaultCharacterTheme
}

func init() {
	bindings := map[string]*input.CommandBinding{}
	bindings["up"] = input.MakeBinding("move-point-up-line")
	bindings["Ctrl-p"] = input.MakeBinding("move-point-up-line")
	bindings["down"] = input.MakeBinding("move-point-down-line")
	bindings["Ctrl-n"] = input.MakeBinding("move-point-down-line")
	bindings["return"] = input.MakeBinding("insert-text", "\n")
	bindings["pageup"] = input.MakeBinding("scroll-page-up")
	bindings["pagedown"] = input.MakeBinding("scroll-page-down")

	input.AddInsertCommands(bindings, "insert-text")
	input.AddSingleLineEditCommands(bindings)
	interfaces.AddMode("findfile", findFileModeFactory{})
}
