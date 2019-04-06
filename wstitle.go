package wstitle

import (
	"log"
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
  Name string
  Number string
  Sep string
  Title string
}

func GetWsName() (wsName, error) {
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

func GetNewNameWindow(ws wsName) (str string) {
	str, ok, err := dlgs.Entry("wstitle", "Set workspace title", ws.Title)
	if !ok {
		log.Fatalln(err)
	}
  return str
}

func GetNewNameDmenu(name wsName) (str string) {
  var leafList []*i3.Node
  ws := getCurrentWorkspace()
  leaves := walkTree(ws, leafList)
  // for _, node := range children {
  //   fmt.Println("result:", node.Name)
  // }
  inList := []string{name.Title}
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
