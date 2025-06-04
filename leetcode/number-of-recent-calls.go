type RecentCounter struct {
	Requests []int
}

func Constructor() RecentCounter {
	return RecentCounter{Requests: []int{}}
}

func binarySearch(x int, array []int) int {
	l, r := 0, len(array)-1
	for l <= r {
		m := (l + r) / 2
		switch {
		case array[m] == x:
			return m
		case array[m] < x:
			l = m + 1
		default:
			r = m - 1
		}
	}
	return l
}

func (this *RecentCounter) Ping(t int) int {
	this.Requests = append(this.Requests, t)
	p := binarySearch(t-3000, this.Requests)
	this.Requests = this.Requests[p:]
	return len(this.Requests)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
