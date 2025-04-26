// DisjointSet represents the data structure
type DisjointSet struct {
	parent []byte
	rank   []byte
}

// NewDisjointSet initializes a new disjoint set
func NewDisjointSet(size int) *DisjointSet {
	ds := &DisjointSet{
		parent: make([]byte, size),
		rank:   make([]byte, size),
	}
	for i := range ds.parent {
		ds.parent[i] = byte(i) // Cast int to byte
	}
	return ds
}

// Find returns the representative of the set containing x
// Implements path compression
func (ds *DisjointSet) Find(x byte) byte {
	if ds.parent[x] != x {
		ds.parent[x] = ds.Find(ds.parent[x]) // Path compression
	}
	return ds.parent[x]
}

// Union merges the sets containing x and y
// Uses union by rank optimization
func (ds *DisjointSet) Union(x, y byte) {
	rootX := ds.Find(x)
	rootY := ds.Find(y)

	if rootX != rootY {
		if ds.rank[rootX] < ds.rank[rootY] {
			ds.parent[rootX] = rootY
		} else if ds.rank[rootX] > ds.rank[rootY] {
			ds.parent[rootY] = rootX
		} else {
			ds.parent[rootY] = rootX
			ds.rank[rootX]++
		}
	}
}

func buildDisjoinSet(s1 string, s2 string) *DisjointSet {
	ds := NewDisjointSet(26)
	n := len(s1)
	for i := 0; i < n; i++ {
		x, y := s1[i]-'a', s2[i]-'a'
		ds.Union(x, y)
	}
	return ds
}

func getMinimals(ds *DisjointSet) map[byte]byte {
	minimals := map[byte]byte{}
	for x := byte(0); x < 26; x++ {
		parent := ds.Find(x)
		_, ok := minimals[parent]
		if !ok {
			minimals[parent] = x
		}
		minimals[parent] = min(minimals[parent], x)
	}
	return minimals
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	ds := buildDisjoinSet(s1, s2)
	minimals := getMinimals(ds)
	n := len(baseStr)
	output := make([]byte, n)
	for i := 0; i < n; i++ {
		x := ds.Find(baseStr[i] - 'a')
		output[i] = minimals[x] + 'a'
	}
	return string(output)
}
