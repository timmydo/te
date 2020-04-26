package modes

import (
	"github.com/timmydo/te/input"
	"github.com/timmydo/te/interfaces"
	"github.com/timmydo/te/theme"
)

type findFileModeFactory struct {
}

func (f findFileModeFactory) Create() interfaces.EditorMode {
	return &findFileMode{}
}

type findFileMode struct {
	selectedLine int
}

func (m findFileMode) Name() string {
	return "findfile"
}

func (m findFileMode) ExecuteCommand(w interfaces.Window, key *input.KeyPressInfo) error {
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
	interfaces.AddMode("findfile", findFileModeFactory{})
}
