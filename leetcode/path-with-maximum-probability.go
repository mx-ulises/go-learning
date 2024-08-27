type Edge struct {
	Node int
	Prob float64
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Edge

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Prob > pq[j].Prob
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Edge)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return item
}

func BuildGraph(edges *[][]int, succProb *[]float64) *map[int][]Edge {
	graph := map[int][]Edge{}
	for i := 0; i < len(*edges); i++ {
		a, b := (*edges)[i][0], (*edges)[i][1]
		p := (*succProb)[i]
		graph[a] = append(graph[a], Edge{Node: b, Prob: p})
		graph[b] = append(graph[b], Edge{Node: a, Prob: p})
	}
	return &graph
}

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	graph := BuildGraph(&edges, &succProb)
	h := &PriorityQueue{&Edge{Node: start_node, Prob: 1.0}}
	heap.Init(h)
	visited := map[int]bool{}
	for 0 < h.Len() {
		current := heap.Pop(h).(*Edge)
		node, prob := current.Node, current.Prob
		if node == end_node {
			return prob
		}
		if visited[node] {
			continue
		}
		visited[node] = true
		for _, succ := range (*graph)[node] {
			heap.Push(h, &Edge{Node: succ.Node, Prob: succ.Prob * prob})
		}
	}
	return 0.0
}
