package ui

import (
	"fmt"
	"log"
	"os"
	"unicode/utf8"

	"github.com/timmydo/te/buffer"
	"github.com/timmydo/te/commands"
	"github.com/timmydo/te/input"
	"github.com/timmydo/te/theme"
	"github.com/timmydo/te/widgets"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func setColor(cr *cairo.Context, c theme.Color) {
	cr.SetSourceRGBA(c.R, c.G, c.B, c.A)
}

func drawPanel(win *widgets.Window, cr *cairo.Context, x, y, width, height float64) {
	setColor(cr, theme.LeftPanelBackgroundColor)
	cr.Rectangle(0, 0, width, height)
	cr.Fill()
}

func getLineToPrint(lineBytes []byte, runesPrinted int, runesOnLine int) string {
	byteOffset := buffer.RuneToByteIndex(runesPrinted, lineBytes)
	byteEnd := buffer.RuneToByteIndex(runesPrinted+runesOnLine, lineBytes)
	return string(lineBytes[byteOffset:byteEnd])
}

func drawEditors(win *widgets.Window, cr *cairo.Context, x, y, width, height float64) {
	cr.SelectFontFace("monospace", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
	fontSize := 14.0
	cr.SetFontSize(fontSize)
	//	log.Printf("drawEditors %v\n", win.OpenBuffer)
	if win.OpenBuffer != nil {

		loc := win.OpenBuffer.GetScrollPosition()
		lines := win.OpenBuffer.Data.Contents
		ypos := y + fontSize
		line := loc.Y
		lineEnd := lines.End().Y + 1

		lineNumberExtents := cr.TextExtents(fmt.Sprintf(" %d", (line+1)*10))
		characterExtents := cr.TextExtents("X")
		point := win.OpenBuffer.Point

		setColor(cr, theme.LineNumberBackgroundColor)
		cr.Rectangle(x, y, lineNumberExtents.XAdvance, height)
		cr.Fill()
		// log.Printf("fill: %v %v %v %v\n", x, y, x+lineNumberExtents.XAdvance, height)
		// log.Printf("lne: %v\n", lineNumberExtents)
		runeWidth := characterExtents.XAdvance
		runeHeight := fontSize
		textOffsetFromLineNumberColumn := characterExtents.XBearing
		textStartX := x + lineNumberExtents.XAdvance + textOffsetFromLineNumberColumn
		runesPerLine := int((width - x) / runeWidth)
		for ypos < height && line < lineEnd {
			lineBytes := lines.LineBytes(line)

			// print line number
			cr.MoveTo(x, ypos)
			setColor(cr, theme.LineNumberFontColor)
			cr.ShowText(fmt.Sprintf(" %d", line+1))
			runesOnLine := utf8.RuneCount(lineBytes)

			// print cursor
			if line == point.Y {
				setColor(cr, theme.CursorColor)
				pointXPos := point.X
				if pointXPos > runesOnLine {
					pointXPos = runesOnLine
				}

				// log.Printf("cursor %v %v %v %v\n", textStartX+(float64(pointXPos)*runeWidth), ypos, runeWidth, runeHeight)
				cr.Rectangle(textStartX+(float64(pointXPos)*runeWidth),
					ypos+characterExtents.YBearing,
					runeWidth,
					runeHeight)
				cr.Fill()
				setColor(cr, theme.PrimaryFontColor)
			}

			for runesPrinted := 0; runesPrinted < runesOnLine; {
				// log.Printf("@ (%v, %v) +%d: %s\n", x, ypos, runesPrinted, string(lineBytes))

				// print text on line
				cr.MoveTo(textStartX, ypos)
				setColor(cr, theme.PrimaryFontColor)
				cr.ShowText(getLineToPrint(lineBytes, runesPrinted, runesOnLine))
				runesPrinted += runesPerLine
			}

			ypos += fontSize
			line++

		}
	}
}

func draw(win *widgets.Window, da *gtk.DrawingArea, cr *cairo.Context) {
	target := cr.GetTarget()
	height := float64(target.GetHeight())
	width := float64(target.GetWidth())
	leftCol := width * win.LeftPanelWidthPercent / 100.0

	// log.Printf("draw(%v) size %v x %v\n", win, width, height)

	drawPanel(win, cr, 0, 0, leftCol, height)
	drawEditors(win, cr, leftCol, 0, width, height)
}

func keyPressEvent(teW *widgets.Window, win *gtk.Window, ev *gdk.Event) {
	keyEvent := &gdk.EventKey{ev}
	item, found := keyMap[keyEvent.KeyVal()]
	if found {
		log.Printf("Handle keypress %v\n", item)
		cmd := input.FindCommand(item)
		if cmd != nil {

			err := commands.GlobalCommands.ExecuteCommand(teW, cmd)
			if err != nil {
				log.Printf("Error: %v\n", err.Error())
			}
			win.QueueDraw()
		}
	} else {
		log.Printf("Key not found %d\n", keyEvent.KeyVal())
	}

}

func setupWindow(teW *widgets.Window, win *gtk.Window) {
	win.SetTitle("TE")
	win.Connect("destroy", func() {
		widgets.ApplicationInstance.KillWindow(teW)
		if len(widgets.ApplicationInstance.Windows) == 0 {
			gtk.MainQuit()
		}
	})

	// Create a new label widget to show in the window.
	da, err := gtk.DrawingAreaNew()
	if err != nil {
		log.Fatal("Unable to create drawing area:", err)
	}

	da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) { draw(teW, da, cr) })
	win.Connect("key-press-event", func(win *gtk.Window, ev *gdk.Event) { keyPressEvent(teW, win, ev) })

	// Add the label to the window.
	win.Add(da)

	// Set the default window size.
	win.SetDefaultSize(800, 600)

	// Recursively show all widgets contained in this window.
	win.ShowAll()
}

func Start() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to get working directory:", err)
	}
	teWindow := widgets.ApplicationInstance.CreateWindow("te", cwd)
	setupWindow(teWindow, win)

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}
