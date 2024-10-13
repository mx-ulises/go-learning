// Define the struct
type Element struct {
	Val       int
	ListIndex int
}

// Define a type that implements heap.Interface
type MinHeap []Element

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func smallestRange(nums [][]int) []int {
	h := &MinHeap{}
	listCount, listIndex := map[int]int{}, map[int]int{}
	coveredLists, n := 0, len(nums)
	dq := []Element{}
	output, d := []int{}, 200001
	for i := 0; i < n; i++ {
		heap.Push(h, Element{Val: nums[i][0], ListIndex: i})
		listIndex[i] = 1
	}
	for 0 < h.Len() {
		element := heap.Pop(h).(Element)
		i := element.ListIndex
		if listIndex[i] < len(nums[i]) {
			j := listIndex[i]
			newElement := Element{Val: nums[i][j], ListIndex: i}
			heap.Push(h, newElement)
			listIndex[i]++
		}
		dq = append(dq, element)
		if listCount[i] == 0 {
			coveredLists++
		}
		listCount[i]++
		for coveredLists == n {
			l, r := dq[0].Val, dq[len(dq)-1].Val
			diff := r - l
			if diff < d {
				output = []int{l, r}
				d = diff
			}
			i = dq[0].ListIndex
			dq = dq[1:]
			listCount[i]--
			if listCount[i] == 0 {
				coveredLists--
			}
		}
	}
	return output
}
