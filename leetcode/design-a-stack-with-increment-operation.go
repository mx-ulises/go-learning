type CustomStack struct {
	Storage  []int
	Capacity int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{Capacity: maxSize}
}

func (this *CustomStack) Push(x int) {
	if len(this.Storage) == this.Capacity {
		return
	}
	this.Storage = append(this.Storage, x)
}

func (this *CustomStack) Pop() int {
	n := len(this.Storage)
	if n == 0 {
		return -1
	}
	x := this.Storage[n-1]
	this.Storage = this.Storage[:n-1]
	return x
}

func (this *CustomStack) Increment(k int, val int) {
	k = min(k, len(this.Storage))
	for i := 0; i < k; i++ {
		this.Storage[i] += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
