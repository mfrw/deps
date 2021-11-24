package graph

func (g *Graph) TopologicalSort(u int) []int {
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

	if cycleFound {
		return nil
	}
	return result
}
