package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joepvd/wstitle"
	"go.i3wm.org/i3"
)

func main() {
	var mode, dmenuCommand string
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "%s is a workspace rename utility for i3wm and sway\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.StringVar(&mode, "mode", "window", "how to select (possible values: window, dmenu")
	flag.StringVar(&dmenuCommand, "dmenu", "dmenu-run", "The dmenu command")
	flag.Parse()

	var ok bool
	var name wstitle.WsName
	var str string

	ws, err := wstitle.ActiveWorkspace()
	if err != nil {
		log.Fatalf("Crap")
	}

	switch mode {
	case "window":
		name = ws
	case "dmenu":
		name, ok = wstitle.ActiveWindow()
		if !ok {
			log.Fatalf("It is bad, mkay?\n")
		}
	default:
		log.Fatalf("Do not understand Mode %s\n", mode)
	}
	str = wstitle.Ask(name.Title)

	newTitle := fmt.Sprintf("%s%s%s", ws.Number, ws.Sep, str)
	_, err = i3.RunCommand(fmt.Sprintf(`rename workspace "%s" to "%s"`, ws.Name, newTitle))
	if err != nil {
		log.Fatalln(err)
	}
}
