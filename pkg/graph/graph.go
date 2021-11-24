package graph

import "sync"

type Graph struct {
	mu      sync.RWMutex
	AdjList map[int]map[int]struct{}
}

type Edge struct {
	From int
	To   int
}

func NewGraph() *Graph {
	return &Graph{
		AdjList: make(map[int]map[int]struct{}),
	}
}

func (g *Graph) AddNode(n int) {
	g.mu.Lock()
	g.AdjList[n] = make(map[int]struct{})
	g.mu.Unlock()
}

func (g *Graph) AddEdge(from, to int) {
	if _, ok := g.AdjList[from]; !ok {
		g.AddNode(from)
	}
	if _, ok := g.AdjList[to]; !ok {
		g.AddNode(to)
	}

	g.mu.Lock()
	g.AdjList[from][to] = struct{}{}
	g.mu.Unlock()
}

func (g *Graph) GetAllEdgesFromChan(from int) <-chan int {
	ch := make(chan int)
	go func() {
		g.mu.RLock()
		if vv, ok := g.AdjList[from]; ok {
			for v := range vv {
				ch <- v
			}
		}
		close(ch)
		g.mu.RUnlock()
	}()
	return ch
}

func (g *Graph) GetAllEdgesChan() <-chan *Edge {
	ch := make(chan *Edge)
	go func() {
		g.mu.RLock()
		for k, vv := range g.AdjList {
			for v := range vv {
				ch <- &Edge{k, v}
			}
		}
		close(ch)
		g.mu.RUnlock()
	}()
	return ch
}

func (g *Graph) FindCycle(u int) bool {
	g.mu.RLock()
	defer g.mu.RUnlock()

	color := make([]uint8, len(g.AdjList), len(g.AdjList))
	result := []int{}

	cycleFound := false

	var dfs func(u int)
	dfs = func(u int) {
		if color[u] == 2 {
			return
		}
		color[u] = 1

		for v := range g.AdjList[u] {
			if color[v] == 0 {
				dfs(v)
			} else if color[v] == 1 {
				cycleFound = true
				return
			}
		}
		result = append(result, u)
		color[u] = 2

	}

	dfs(u)

	return cycleFound
}
