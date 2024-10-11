// Define the struct
type Seat struct {
	LeavingTime int
	Chair       int
}

// Define a type that implements heap.Interface and holds Seats
type SeatHeap []Seat

func (h SeatHeap) Len() int           { return len(h) }
func (h SeatHeap) Less(i, j int) bool { return h[i].LeavingTime < h[j].LeavingTime }
func (h SeatHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push and Pop use pointer receivers because they modify the slice's length
func (h *SeatHeap) Push(x interface{}) {
	*h = append(*h, x.(Seat))
}

func (h *SeatHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push and Pop use pointer receivers because they modify the slice's length
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func smallestChair(times [][]int, targetFriend int) int {
	tartgetFriendArrivad := times[targetFriend][0]
	sort.Slice(times, func(i, j int) bool {
		return times[i][0] < times[j][0]
	})
	seats := &SeatHeap{}
	available := &IntHeap{}
	for _, person := range times {
		arrival, leaving := person[0], person[1]
		for 0 < seats.Len() && (*seats)[0].LeavingTime <= arrival {
			person := heap.Pop(seats).(Seat)
			heap.Push(available, person.Chair)
		}
		chair := seats.Len()
		if 0 < available.Len() {
			chair = heap.Pop(available).(int)
		}
		heap.Push(seats, Seat{LeavingTime: leaving, Chair: chair})
		if arrival == tartgetFriendArrivad {
			return chair
		}
	}
	return -1
}
