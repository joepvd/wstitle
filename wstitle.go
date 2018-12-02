package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/gen2brain/dlgs"
	"go.i3wm.org/i3"
)

func help() {
	if len(os.Args) > 1 {
		arg := os.Args[1]
		helpText := fmt.Sprintf(`
"%s" is a workspace rename utility for i3wm and sway.
Valid command line options:
  "-h" or "-help" or "--help"
`, os.Args[0])
		if arg == "-h" || arg == "--help" || arg == "-help" {
			fmt.Println(helpText)
			os.Exit(0)
		} else {
			os.Stderr.WriteString(helpText)
			os.Exit(1)
		}
	}
}

func getCurrentWorkspace() (ws *i3.Node) {
	tree, err := i3.GetTree()
	if err != nil {
		log.Fatalln(err)
	}
	ws = tree.Root.FindFocused(func(n *i3.Node) bool {
		return n.Type == i3.WorkspaceNode
	})
	if ws == nil {
		log.Fatalln("could not locate workspace")
	}
	return
}

func getReParams(regEx, str string) (reMap map[string]string) {
	r := regexp.MustCompile(regEx)
	match := r.FindStringSubmatch(str)
	reMap = make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i > 0 && i <= len(match) {
			reMap[name] = match[i]
		}
	}
	return
}

func main() {
	help()
	ws := getCurrentWorkspace()
	curWsTitle := getReParams(`^((?P<Number>\d+)(?P<Sep>: ))?(?P<Title>.*)`, ws.Name)
	var number, sep, title string
	if curWsTitle["Number"] == "" {
		number = curWsTitle["Title"]
		sep = ": "
		title = ""
	} else {
		title = curWsTitle["Title"]
		number = curWsTitle["Number"]
		sep = curWsTitle["Sep"]
	}

	str, ok, err := dlgs.Entry("wstitle", "Set workspace title", title)
	if !ok {
		log.Fatalln(err)
	}
	newTitle := fmt.Sprintf("%s%s%s", number, sep, str)
	_, err = i3.RunCommand(fmt.Sprintf(`rename workspace "%s" to "%s"`, ws.Name, newTitle))
	if err != nil {
		log.Fatalln(err)
	}
}
