package pkg

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/mzohreva/GoGraphviz/graphviz"
)

type Graph struct {
	graphviz.Graph
}

func (g *Graph) OutputDOT(writer io.Writer) {
	g.Graph.GenerateDOT(writer)
}

func NewGraph(modules map[string]Module) *Graph {
	g := graphviz.Graph{}
	nodes := make(map[string]int, len(modules))
	for name, module := range modules {
		moduleType, _ := strings.CutPrefix(reflect.TypeOf(module).String(), "*pkg.")
		nodes[name] = g.AddNode(fmt.Sprintf("%s\n%s", name, moduleType))
	}
	for name, module := range modules {
		for _, dest := range module.Destinations() {
			id := g.AddEdge(nodes[name], nodes[dest], "dest")
			g.EdgeAttribute(id, "dir", "forward")
		}
	}
	return &Graph{g}
}
