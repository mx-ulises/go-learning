type OrderedStream struct {
	Memory  []string
	Pointer int
	Size    int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{Memory: make([]string, n), Size: n, Pointer: 0}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	output := []string{}
	this.Memory[idKey-1] = value
	for this.Pointer < this.Size && this.Memory[this.Pointer] != "" {
		output = append(output, this.Memory[this.Pointer])
		this.Pointer++
	}
	return output
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
