package dsu

type DSU struct {
	parent []int
	rank   []int
	cmps   int
}

func NewDSU(sz int) *DSU {
	par := make([]int, sz, sz)
	rnk := make([]int, sz, sz)
	for i := 0; i < sz; i++ {
		par[i] = i
		rnk[i] = 0
	}
	return &DSU{
		parent: par,
		rank:   rnk,
		cmps:   sz,
	}
}

func (d *DSU) Unite(a, b int) {
	a, b = d.Find(a), d.Find(b)
	if a != b {
		if d.rank[a] < d.rank[b] {
			a, b = b, a
		}
		d.parent[b] = a
		d.cmps--
		if d.rank[a] == d.rank[b] {
			d.rank[a]++
		}
	}
}

func (d *DSU) Find(u int) int {
	if u != d.parent[u] {
		d.parent[u] = d.Find(d.parent[u])
	}
	return d.parent[u]
}

func (d *DSU) Components() int {
	return d.cmps
}
