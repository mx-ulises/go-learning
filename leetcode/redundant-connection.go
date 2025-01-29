type DisjointSet struct {
	parent []int
	rank   []int
}

func NewDisjointSet(n int) *DisjointSet {
	ds := &DisjointSet{
		parent: make([]int, n),
		rank:   make([]int, n),
	}
	for i := range ds.parent {
		ds.parent[i] = i
	}
	return ds
}

func (ds *DisjointSet) Find(u int) int {
	if ds.parent[u] != u {
		ds.parent[u] = ds.Find(ds.parent[u])
	}
	return ds.parent[u]
}

func (ds *DisjointSet) Union(u, v int) {
	rootU := ds.Find(u)
	rootV := ds.Find(v)

	if rootU != rootV {
		if ds.rank[rootU] > ds.rank[rootV] {
			ds.parent[rootV] = rootU
		} else if ds.rank[rootU] < ds.rank[rootV] {
			ds.parent[rootU] = rootV
		} else {
			ds.parent[rootV] = rootU
			ds.rank[rootU]++
		}
	}
}

func findRedundantConnection(edges [][]int) []int {
	ds := NewDisjointSet(len(edges) + 1)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		if ds.Find(a) == ds.Find(b) {
			return edge
		}
		ds.Union(a, b)
	}
	return nil
}
