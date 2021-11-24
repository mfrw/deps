package main

import (
	"flag"
	"fmt"

	"github.com/mfrw/deps/pkg/dsu"
)

var (
	input = flag.String("i", "/home/mfrw/mariner-org/parity/build/pkg_artifacts/graph.dot", "Input Dot Graph file")
)

func main() {
	flag.Parse()
	graphInput := *input

	ng, _ := NewNodeGraphFromDotGraphFile(graphInput)

	fmt.Println("NR Items in graph:", len(ng.nm))
	fmt.Println("NR Connected cmps:", DistinctComponents(ng))
	fmt.Println("NR Node Types:", NodeTypes(ng))
	fmt.Println("Cycle Found:", ng.FindCycle())
}

func DistinctComponents(ng *NodeGraph) int {
	d := dsu.NewDSU(ng.Len())
	for edge := range ng.GetAllEdgesChan() {
		d.Unite(edge.From, edge.To)
	}
	return d.Components()
}

func NodeTypes(ng *NodeGraph) map[string]int {
	mp := map[string]int{}

	for _, v := range ng.nm {
		mp[v.Type]++
	}
	return mp
}
