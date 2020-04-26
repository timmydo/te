package theme

type Color struct {
	R float64
	G float64
	B float64
	A float64
}

type BufferThemeStyle struct {
	LineNumberFont       Color
	LineNumberBackground Color
	Background           Color
}

type LineThemeStyle struct {
	Background Color
}

type CharacterThemeStyle struct {
	Font      Color
	Cursor    Color
	Selection Color
}

var (
	DefaultBufferTheme = &BufferThemeStyle{
		LineNumberFont:       Color{0.0, 0.0, 0.0, 1.0},
		LineNumberBackground: Color{0.9, 0.9, 0.9, 1.0},
		Background:           Color{1.0, 1.0, 1.0, 1.0},
	}
	DefaultLineTheme = &LineThemeStyle{
		Background: Color{0.95, 0.95, 0.95, 1.0},
	}
	DefaultCharacterTheme = &CharacterThemeStyle{
		Font:      Color{0.0, 0.0, 0.0, 1.0},
		Cursor:    Color{0.9, 0.4, 0.4, 1.0},
		Selection: Color{0.9, 0.9, 0.4, 1.0},
	}
)
