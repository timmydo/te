package main

import (
	"flag"
	"fmt"

	"github.com/timmydo/te/ui"
)

var (
	// Command line flags
	flagVersion = flag.Bool("version", false, "Show the version number and information")
	flagDebug   = flag.Bool("debug", false, "Enable debug mode (prints debug info to ./log.txt)")
	optionFlags map[string]*string
	Version     string
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

	ui.Start()

}
