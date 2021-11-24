package main

import (
	"flag"
	"fmt"

	"github.com/mfrw/deps/pkg/graph"
	"github.com/mfrw/deps/pkg/node"
)

var (
	input = flag.String("i", "./graph.dot", "Input Dot Graph file")
)

type NodeGraph struct {
	g  *graph.Graph
	nm map[int]*node.Node
}

func NewNodeGraph() *NodeGraph {
	return &NodeGraph{
		g:  graph.NewGraph(),
		nm: make(map[int]*node.Node),
	}
}

func NewNodeGraphFromDotGraphFile(fname string) (*NodeGraph, error) {
	edges, err := node.GetOnlyEdges(fname)

	if err != nil {
		return nil, err
	}

	graph := NewNodeGraph()

	for _, edge := range edges {
		from, to, err := node.ProcessLineToNode(edge)
		if err != nil {
			return nil, err
		}
		graph.AddEdge(from, to)
	}
	return graph, nil
}

func (ng *NodeGraph) AddNode(n *node.Node) {
	ng.g.AddNode(n.Id)
	ng.nm[n.Id] = n
}

func (ng *NodeGraph) AddEdge(from, to *node.Node) {
	ng.g.AddEdge(from.Id, to.Id)
	ng.nm[from.Id] = from
	ng.nm[to.Id] = to
}

func main() {
	flag.Parse()
	graphInput := *input

	ng, _ := NewNodeGraphFromDotGraphFile(graphInput)

	fmt.Println("NR Items in graph:", len(ng.nm))
}
