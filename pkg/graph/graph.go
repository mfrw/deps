package graph

type Graph struct {
	AdjList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		AdjList: make(map[int][]int),
	}
}

func (g *Graph) AddNode(n int) {
	g.AdjList[n] = make([]int, 0)
}

func (g *Graph) AddEdge(from, to int) {
	if _, ok := g.AdjList[from]; !ok {
		g.AddNode(from)
	}
	if _, ok := g.AdjList[to]; !ok {
		g.AddNode(to)
	}
	g.AdjList[from] = append(g.AdjList[from], to)
}
