type Element struct {
	Count int
	Char  rune
}

type MaxHeap []Element

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].Count > h[j].Count }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func AddIfValid(h *MaxHeap, element *Element) {
	if 0 < element.Count {
		heap.Push(h, *element)
	}
}

func longestDiverseString(a int, b int, c int) string {
	h := &MaxHeap{}
	AddIfValid(h, &Element{Count: a, Char: 'a'})
	AddIfValid(h, &Element{Count: b, Char: 'b'})
	AddIfValid(h, &Element{Count: c, Char: 'c'})
	output := []rune{}
	n := 0
	for 0 < h.Len() {
		element := heap.Pop(h).(Element)
		if 2 <= n && output[n-1] == element.Char && output[n-2] == element.Char {
			if h.Len() == 0 {
				break
			}
			aux := heap.Pop(h).(Element)
			heap.Push(h, element)
			element = aux
		}
		output = append(output, element.Char)
		n++
		element.Count--
		AddIfValid(h, &element)
	}
	return string(output)
}
