type MyCircularDeque struct {
	Storage  []int
	Front    int
	Back     int
	Size     int
	Capacity int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{Storage: make([]int, k), Capacity: k}
}

func (this *MyCircularDeque) DecreaseFront() {
	this.Front -= 1
	if this.Front == -1 {
		this.Front = this.Capacity - 1
	}
}

func (this *MyCircularDeque) IncreaseFront() {
	this.Front += 1
	if this.Front == this.Capacity {
		this.Front = 0
	}
}

func (this *MyCircularDeque) DecreaseBack() {
	this.Back -= 1
	if this.Back == -1 {
		this.Back = this.Capacity - 1
	}
}

func (this *MyCircularDeque) IncreaseBack() {
	this.Back += 1
	if this.Back == this.Capacity {
		this.Back = 0
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.Size == this.Capacity {
		return false
	}
	this.Storage[this.Front] = value
	this.DecreaseFront()
	this.Size++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.Size == this.Capacity {
		return false
	}
	this.IncreaseBack()
	this.Storage[this.Back] = value
	this.Size++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.Size == 0 {
		return false
	}
	this.IncreaseFront()
	this.Storage[this.Front] = 0
	this.Size--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.Size == 0 {
		return false
	}
	this.Storage[this.Back] = 0
	this.DecreaseBack()
	this.Size--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.Size == 0 {
		return -1
	}
	position := this.Front + 1
	if position == this.Capacity {
		position = 0
	}
	return this.Storage[position]
}

func (this *MyCircularDeque) GetRear() int {
	if this.Size == 0 {
		return -1
	}
	return this.Storage[this.Back]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.Size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.Size == this.Capacity
}
