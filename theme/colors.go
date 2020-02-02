package theme

type Color struct {
	R float64
	G float64
	B float64
	A float64
}

var (
	LeftPanelBackgroundColor  = Color{0.7, 0.7, 0.7, 1.0}
	PrimaryFontColor          = Color{0.0, 0.0, 0.0, 1.0}
	LineNumberFontColor       = Color{0.0, 0.0, 0.0, 1.0}
	LineNumberBackgroundColor = Color{0.9, 0.9, 0.9, 1.0}
	CursorColor               = Color{0.9, 0.4, 0.4, 1.0}
)
