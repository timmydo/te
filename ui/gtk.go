package ui

import (
	"fmt"
	"log"
	"os"

	"github.com/timmydo/te/buffer"
	"github.com/timmydo/te/input"
	"github.com/timmydo/te/widgets"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

var (
	keyMap map[uint]*input.KeyPressInfo
)

func init() {
	keyMap = make(map[uint]*input.KeyPressInfo)
	keyMap[0x61] = input.NewKeyPressInfo("a")

}

func drawPanel(win *widgets.Window, cr *cairo.Context, x, y, width, height float64) {
	cr.SetSourceRGB(.7, .7, .7)
	cr.Rectangle(0, 0, width, height)
	cr.Fill()
}

func drawEditors(win *widgets.Window, cr *cairo.Context, x, y, width, height float64) {
	cr.SelectFontFace("DejaVuSansMono", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
	fontSize := 14.0
	cr.SetFontSize(fontSize)
	cr.SetSourceRGB(0, 0, 0)
	log.Printf("drawEditors %v\n", win.OpenBuffer)
	if win.OpenBuffer != nil {

		loc := win.OpenBuffer.GetScrollPosition()
		lines := win.OpenBuffer.Data.Contents
		ypos := y + fontSize
		line, col := loc.Y, loc.X
		lineEnd := lines.End().Y

		lineNumberExtents := cr.TextExtents(fmt.Sprintf(" %d:", (line+1)*10))

		for ypos < height && line < lineEnd {

			lineBytes := lines.LineBytes(line)
			startOffset := buffer.RuneToByteIndex(col, lineBytes)
			log.Printf("@ (%v, %v) +%d: %s\n", x, ypos, startOffset, string(lineBytes))
			if startOffset < len(lineBytes) {
				cr.MoveTo(x, ypos)
				cr.ShowText(fmt.Sprintf(" %d:", line+1))
				cr.MoveTo(x+lineNumberExtents.XAdvance, ypos)
				cr.ShowText(string(lineBytes[startOffset:]))
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

	log.Printf("draw(%v) size %v x %v\n", win, width, height)

	drawPanel(win, cr, 0, 0, leftCol, height)
	drawEditors(win, cr, leftCol, 0, width, height)
}

func keyPressEvent(teW *widgets.Window, win *gtk.Window, ev *gdk.Event) {
	keyEvent := &gdk.EventKey{ev}
	item, found := keyMap[keyEvent.KeyVal()]
	if found {
		teWindow := widgets.ApplicationInstance.FindWindow(teW)
		log.Println("Handle keypress %v\n", item)
		teWindow.HandleKeyPress(item)
		win.QueueDraw()
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
