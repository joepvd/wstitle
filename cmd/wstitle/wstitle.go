package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"go.i3wm.org/i3"

	"github.com/joepvd/wstitle"
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

	ws, err := wstitle.GetWsName()

	var str string
	switch mode {
	case "window":
		str = wstitle.GetNewNameWindow(ws)
	case "dmenu":
		str = wstitle.GetNewNameDmenu(ws)
	default:
		log.Fatalf("Do not understand Mode %s\n", mode)
	}

	newTitle := fmt.Sprintf("%s%s%s", ws.Number, ws.Sep, str)
	_, err = i3.RunCommand(fmt.Sprintf(`rename workspace "%s" to "%s"`, ws.Name, newTitle))
	if err != nil {
		log.Fatalln(err)
	}
}
