package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

var (
	// Command line flags
	flagVersion = flag.Bool("version", false, "Show the version number and information")
	flagDebug   = flag.Bool("debug", false, "Enable debug mode (prints debug info to ./log.txt)")
	optionFlags map[string]*string

	Version string
)

func printHelp() {
	fmt.Println("Help")
}

func main() {

	flag.Parse()
	if *flagVersion {
		fmt.Println("te " + Version)
		return
	}

	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("TE")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Create a new label widget to show in the window.
	l, err := gtk.LabelNew("Hello, gotk3!")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

	// Add the label to the window.
	win.Add(l)

	// Set the default window size.
	win.SetDefaultSize(800, 600)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}
