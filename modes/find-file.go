package modes

import (
	"errors"
	"log"

	"github.com/timmydo/te/commands"
	"github.com/timmydo/te/input"
	"github.com/timmydo/te/interfaces"
	"github.com/timmydo/te/theme"
)

type findFileModeFactory struct {
	bindings map[string]*input.CommandBinding
}

func (f findFileModeFactory) Create() interfaces.EditorMode {
	return &findFileMode{f.bindings, 0}
}

type findFileMode struct {
	bindings     map[string]*input.CommandBinding
	selectedLine int
}

func (m findFileMode) Name() string {
	return "findfile"
}

func (m findFileMode) ExecuteCommand(w interfaces.Window, key string) error {
	if err, handled := m.handleModeBinding(w, key); handled {
		return err
	}
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

func (this *findFileMode) handleModeBinding(w interfaces.Window, key string) (error, bool) {

	switch key {
	case "up":
	case "Ctrl-p":
		if this.selectedLine > 0 {
			this.selectedLine--
		}
		return nil, true

	case "down":
	case "Ctrl-n":
		if this.selectedLine < w.OpenBuffer().GetLines().End().Y {
			this.selectedLine++
		}
		return nil, true

	case "return":
		lines := w.OpenBuffer().GetLines()
		if this.selectedLine >= 0 && this.selectedLine < lines.End().Y {
			filename := string(lines.LineBytes(this.selectedLine))
			log.Printf("Filename: %s\n", filename)
			bf := interfaces.GetBufferFactory()
			bf.DeleteBuffer(w.OpenBuffer())
			b, err := bf.CreateBufferFromFile(filename)
			w.SetOpenBuffer(b)
			return err, true
		}

		return errors.New("Could not find file"), true
	}

	return nil, false
}

func init() {
	bindings := map[string]*input.CommandBinding{}

	input.AddInsertCommands(bindings, "insert-text")
	input.AddSingleLineEditCommands(bindings)
	interfaces.AddMode("findfile", findFileModeFactory{bindings})
}
