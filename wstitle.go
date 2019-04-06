package main

import (
  "flag"
	"fmt"
	"log"
	"os"
  "strings"
	"regexp"

	"github.com/gen2brain/dlgs"
	"go.i3wm.org/i3"
)


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

type wsName struct {
  name string
  number string
  sep string
  title string
}

func getWsName() (wsName, error) {
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
  name := wsName{ws.Name, number, sep, title}
  return name, nil
}

func getNewNameWindow(ws wsName) (str string) {
	str, ok, err := dlgs.Entry("wstitle", "Set workspace title", ws.title)
	if !ok {
		log.Fatalln(err)
	}
  return str
}

func getNewNameDmenu(name wsName) (str string) {
  var leafList []*i3.Node
  ws := getCurrentWorkspace()
  leaves := walkTree(ws, leafList)
  // for _, node := range children {
  //   fmt.Println("result:", node.Name)
  // }
  inList := []string{name.title}
  for _, leaf := range leaves {
    inList = append(inList, leaf.Name)
  }
  stdin := strings.Join(inList[:], "\n")

  return stdin
}

func walkTree(node *i3.Node, list []*i3.Node) ([]*i3.Node) {
  if len(node.Nodes) == 0 {
    return append(list, node)
  }
  for _, child := range node.Nodes {
    list = walkTree(child, list)
  }
  return list
}

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

  ws, err := getWsName()

  var str string
  switch mode {
  case "window":
    str = getNewNameWindow(ws)
  case "dmenu":
    str = getNewNameDmenu(ws)
  default:
    log.Fatalf("Do not understand Mode %s\n", mode)
  }

	newTitle := fmt.Sprintf("%s%s%s", ws.number, ws.sep, str)
	_, err = i3.RunCommand(fmt.Sprintf(`rename workspace "%s" to "%s"`, ws.name, newTitle))
	if err != nil {
		log.Fatalln(err)
	}
}
