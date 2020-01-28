package ui

import (
	"log"

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

func draw(da *gtk.DrawingArea, cr *cairo.Context) {
	cr.SetSourceRGB(0, 0, 0)
	cr.Rectangle(10, 20, 30, 40)
	cr.Fill()
}

func keyPressEvent(win *gtk.Window, ev *gdk.Event) {
	keyEvent := &gdk.EventKey{ev}
	// fixme lookup key
	item, found := keyMap[keyEvent.KeyVal()]
	if found {
		teWindow := widgets.ApplicationInstance.FindWindow(win)
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
		widgets.ApplicationInstance.KillWindow(win)
		if len(widgets.ApplicationInstance.Windows) == 0 {
			gtk.MainQuit()
		}
	})

	// Create a new label widget to show in the window.
	da, err := gtk.DrawingAreaNew()
	if err != nil {
		log.Fatal("Unable to create drawing area:", err)
	}

	da.Connect("draw", draw)
	win.Connect("key-press-event", keyPressEvent)

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

	teWindow := widgets.ApplicationInstance.CreateWindow("te", win)
	setupWindow(teWindow, win)

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}
