package interfaces

import (
	"github.com/timmydo/te/input"
	"github.com/timmydo/te/theme"
)

type EditorModeFactory interface {
	Create() EditorMode
}

type EditorMode interface {
	Name() string
	ExecuteCommand(w Window, key *input.KeyPressInfo) error
	GetBufferStyle() *theme.BufferThemeStyle
	GetLineStyle(int) *theme.LineThemeStyle
	GetCharacterStyle(int, int) *theme.CharacterThemeStyle
}

var modes map[string]EditorModeFactory

func GetMode(name string) EditorMode {
	return modes[name].Create()
}

func AddMode(name string, m EditorModeFactory) {
	if modes == nil {
		modes = map[string]EditorModeFactory{}
	}

	modes[name] = m
}
