package ui

import (
	"fmt"
	"log"
	"os"
	"unicode/utf8"

	"github.com/timmydo/te/interfaces"
	"github.com/timmydo/te/linearray"
	"github.com/timmydo/te/theme"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func setColor(cr *cairo.Context, c theme.Color) {
	cr.SetSourceRGBA(c.R, c.G, c.B, c.A)
}

func drawBuffer(buf interfaces.Buffer, cr *cairo.Context, x, y, width, height float64) {
	cr.SelectFontFace("monospace", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
	fontSize := 14.0
	cr.SetFontSize(fontSize)
	//	log.Printf("drawEditors %v\n", buf)
	line := buf.ScrollPosition()
	topLine := line
	lines := buf.GetLines()
	ypos := y + fontSize
	lineEnd := lines.End().Y + 1

	lineNumberExtents := cr.TextExtents(fmt.Sprintf(" %d", (line+1)*10))
	characterExtents := cr.TextExtents("X")
	point := buf.Point()
	mark := buf.Mark()

	setColor(cr, buf.Mode().GetBufferStyle().Background)
	cr.Rectangle(x, y, width, height)
	cr.Fill()

	setColor(cr, buf.Mode().GetBufferStyle().LineNumberBackground)
	cr.Rectangle(x, y, lineNumberExtents.XAdvance, height)
	cr.Fill()
	// log.Printf("character extents: %v\n", characterExtents)
	// log.Printf("fill: %v %v %v %v\n", x, y, x+lineNumberExtents.XAdvance, height)
	// log.Printf("lne: %v\n", lineNumberExtents)
	runeWidth := characterExtents.XAdvance
	runeHeight := fontSize
	textOffsetFromLineNumberColumn := characterExtents.XBearing
	textStartX := x + lineNumberExtents.XAdvance + textOffsetFromLineNumberColumn

	for ypos < height && line < lineEnd {
		lineBytes := lines.LineBytes(line)

		// print line number
		cr.MoveTo(x, ypos)
		setColor(cr, buf.Mode().GetBufferStyle().LineNumberFont)
		cr.ShowText(fmt.Sprintf(" %d", line+1))
		runesOnLine := utf8.RuneCount(lineBytes)

		// print line background color
		setColor(cr, buf.Mode().GetLineStyle(line).Background)
		cr.Rectangle(textStartX, ypos+characterExtents.YBearing, width, runeHeight)
		cr.Fill()

		// print cursor
		if line == point.Y {
			pointXPos := point.X
			if pointXPos > runesOnLine {
				pointXPos = runesOnLine
			}

			setColor(cr, buf.Mode().GetCharacterStyle(line, pointXPos).Cursor)
			cr.Rectangle(textStartX+(float64(pointXPos)*runeWidth),
				ypos+characterExtents.YBearing,
				runeWidth,
				runeHeight)
			cr.Fill()
		}

		lineByteOffset := 0
		currentX := textStartX
		currentY := ypos
		// print text on line
		for runesPrinted := 0; runesPrinted < runesOnLine; runesPrinted++ {
			r, rSize := utf8.DecodeRune(lineBytes[lineByteOffset:])
			currentCharacter := string(r)
			lineByteOffset += rSize
			currentLoc := linearray.Loc{runesPrinted, line}
			inSelection := false
			// if mark active
			if mark.Y != -1 {
				if mark.GreaterEqual(point) {
					if currentLoc.GreaterThan(point) && currentLoc.LessEqual(mark) {
						inSelection = true
					}
				} else {
					if currentLoc.GreaterEqual(mark) && currentLoc.LessThan(point) {
						inSelection = true
					}
				}
			}

			// print selection background
			if inSelection {
				setColor(cr, buf.Mode().GetCharacterStyle(line, runesPrinted).Selection)
				cr.Rectangle(textStartX+(float64(runesPrinted)*runeWidth),
					ypos+characterExtents.YBearing,
					runeWidth,
					runeHeight)
				cr.Fill()
			}

			// print character
			cr.MoveTo(currentX, currentY)
			setColor(cr, buf.Mode().GetCharacterStyle(line, runesPrinted).Font)
			cr.ShowText(currentCharacter)
			currentX += characterExtents.XAdvance
			currentY += characterExtents.YAdvance
		}

		ypos += fontSize
		line++

	}

	buf.SetLinesInDisplay(line - topLine)
}

func draw(win interfaces.Window, da *gtk.DrawingArea, cr *cairo.Context) {
	target := cr.GetTarget()
	height := float64(target.GetHeight())
	width := float64(target.GetWidth())

	buf := win.OpenBuffer()
	// log.Printf("draw(%v) size %v x %v, buf: %p\n", win, width, height, buf)
	if buf != nil {
		drawBuffer(buf, cr, 0, 0, width, height)
	}
}

func keyPressEvent(teW interfaces.Window, win *gtk.Window, ev *gdk.Event) {
	keyEvent := &gdk.EventKey{ev}
	keyState := gdk.ModifierType(keyEvent.State())
	item, found := keyMap[keyEvent.KeyVal()]
	if found {
		item.ShiftMod = keyState&gdk.GDK_SHIFT_MASK != 0
		item.CtrlMod = keyState&gdk.GDK_CONTROL_MASK != 0
		item.MetaMod = keyState&gdk.GDK_MOD1_MASK != 0
		item.SuperMod = keyState&gdk.GDK_SUPER_MASK != 0
		item.HyperMod = keyState&gdk.GDK_HYPER_MASK != 0
		err := teW.OpenBuffer().Mode().ExecuteCommand(teW, item.GetName())
		if err != nil {
			log.Printf("Error: %v\n", err.Error())
		}

		win.QueueDraw()
	} else {
		log.Printf("Key not found %d\n", keyEvent.KeyVal())
	}

}

func setupWindow(teW interfaces.Window, win *gtk.Window) {
	win.SetTitle("TE")
	win.Connect("destroy", func() {
		interfaces.GetApplication().KillWindow(teW)
		if len(interfaces.GetApplication().Windows()) == 0 {
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
	app := interfaces.GetApplication()
	log.Printf("App: %v\n", app)
	teWindow := app.CreateWindow("te", cwd)
	setupWindow(teWindow, win)

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}
