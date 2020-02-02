package ui

import (
	"log"
	"os"

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

func draw(teW *widgets.Window, da *gtk.DrawingArea, cr *cairo.Context) {
	cr.SetSourceRGB(0, 0, 0)
	cr.MoveTo(30, 30)
	target := cr.GetTarget()
	height := target.GetHeight()
	width := target.GetWidth()
	log.Printf("%v %d x %d\n", teW, width, height)
	cr.ShowText("hello")
	//	cr.Fill()
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

	teWindow := widgets.ApplicationInstance.CreateWindow("te", os.Getwd())
	setupWindow(teWindow, win)

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}
