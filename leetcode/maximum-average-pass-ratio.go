type Class struct {
	Improvement float64
	Rate        float64
	Pass        float64
	Total       float64
}

type MaxHeap []Class

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].Improvement > h[j].Improvement }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Class))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	h := &MaxHeap{}
	heap.Init(h)
	averageRatio := float64(0)
	for _, classInput := range classes {
		class := Class{Pass: float64(classInput[0]), Total: float64(classInput[1])}
		class.Rate = class.Pass / class.Total
		class.Improvement = (class.Pass+1)/(class.Total+1) - class.Rate
		heap.Push(h, class)
		averageRatio += class.Rate
	}
	for 0 < extraStudents {
		class := heap.Pop(h).(Class)
		averageRatio += class.Improvement
		class.Pass++
		class.Total++
		class.Rate += class.Improvement
		class.Improvement = (class.Pass+1)/(class.Total+1) - class.Rate
		heap.Push(h, class)
		extraStudents--
	}
	return averageRatio / float64(len(classes))
}
