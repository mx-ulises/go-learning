func getBits(num int) int {
	bits := 0
	for num != 0 {
		bits += num & 1
		num >>= 1
	}
	return bits
}

func sortByBits(arr []int) []int {
	buckets := [16][]int{}
	for _, num := range arr {
		bits := getBits(num)
		buckets[bits] = append(buckets[bits], num)
	}
	output := make([]int, len(arr))
	i := 0
	for _, bucket := range buckets {
		sort.Ints(bucket)
		for _, num := range bucket {
			output[i] = num
			i++
		}
	}
	return output
}
