package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joepvd/wstitle"
)

func main() {
	var (
		mode string
		name string
		ok   bool
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "%s is a workspace rename utility for i3wm and sway\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.StringVar(&mode, "mode", "window", "how to select (possible values: window, workspace")
	flag.Parse()

	ws, err := wstitle.ActiveWorkspace()
	if err != nil {
		log.Fatalf("Crap")
	}

	switch mode {
	case "workspace":
		name = ws.Title
	case "window":
		name, ok = wstitle.ActiveWindow()
		if !ok {
			log.Fatalf("It is bad, mkay?\n")
		}
	default:
		log.Fatalf("Do not understand Mode %s\n", mode)
	}

	newName := wstitle.Ask(name)

	if err = wstitle.SetTitle(newName, ws); err != nil {
		log.Fatalln(err)
	}
}
