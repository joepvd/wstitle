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
		newName   string
		ok        bool
		window    bool
		workspace bool
	)

	flag.BoolVar(&window, "window", true, "Use window title as proposal")
	flag.BoolVar(&workspace, "workspace", false, "Use current workspace title as proposal")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "%s is a workspace rename utility for i3wm and sway\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	ws, err := wstitle.ActiveWorkspace()
	if err != nil {
		log.Fatalf("Crap")
	}
	if workspace {
		newName, ok = TryWorkspace(ws)
	} else if window {
		newName, ok = TryWindow(ws)
	}
	if !ok {
		os.Exit(1)
	}

	if err = wstitle.SetTitle(newName, ws); err != nil {
		log.Fatalln(err)
	}
}

func TryWindow(ws wstitle.WsName) (string, bool) {
	win, ok := wstitle.ActiveWindow()
	if !ok {
		log.Fatalf("It is bad, mkay?\n")
	}
	return wstitle.Ask(win.Title)
}

func TryWorkspace(ws wstitle.WsName) (string, bool) {
	return wstitle.Ask(ws.Title)
}
