package wstitle

import (
	"fmt"
	"log"
	"regexp"

	"github.com/gen2brain/dlgs"
	"go.i3wm.org/i3"
)

type WsName struct {
	node   *i3.Node
	Title  string
	Name   string
	Number string
	Sep    string
}

func ActiveWorkspace() (WsName, error) {
	ws := currentWorkspace()
	curWsTitle := getReParams(`^((?P<Number>\d+)(?P<Sep>: ))?(?P<Title>.+)`, ws.Name)
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
	name := WsName{ws, title, ws.Name, number, sep}
	return name, nil
}

func currentWorkspace() (ws *i3.Node) {
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

func ActiveWindow() (WsName, bool) {
	var leafList []*i3.Node
	ws := currentWorkspace()
	for _, leaf := range Leaves(ws, leafList) {
		if leaf.Focused {
			return WsName{leaf, leaf.Name, "", "", ""}, true
		}
	}
	return WsName{}, false
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

func Leaves(node *i3.Node, list []*i3.Node) []*i3.Node {
	if len(node.Nodes) == 0 && node.Type != i3.WorkspaceNode {
		return append(list, node)
	}
	for _, child := range node.Nodes {
		list = Leaves(child, list)
	}
	return list
}

func Ask(name string) (str string, ret bool) {
	str, ok, err := dlgs.Entry("wstitle", "Set workspace title", name)
	if err != nil {
		log.Fatal(err)
	}
	if !ok {
		log.Fatal("cancel pressed")
	}
	if name != str {
		ret = true
	}
	return
}

// XXX: Maybe should be a Write method?
// Consider ReadWriter implementation?
func (ws *WsName) SetTitle(name string) (err error) {
	newTitle := fmt.Sprintf("%s%s%s", ws.Number, ws.Sep, name)
	_, err = i3.RunCommand(fmt.Sprintf(`rename workspace "%s" to "%s"`, ws.Name, newTitle))
	return
}
