func numMod(num, divisor int) int {
	if divisor <= num {
		num -= divisor
	}
	return num
}

func numOfSubarrays(arr []int) int {
	divisor := 1_000_000_007
	sum, current := 0, 0
	prev := map[int]int{0: 1, 1: 0}
	for _, num := range arr {
		current = (current + num) & 1
		if current == 1 {
			sum = numMod(sum+prev[0], divisor)
			prev[1] = numMod(prev[1]+1, divisor)
		} else {
			sum = numMod(sum+prev[1], divisor)
			prev[0] = numMod(prev[0]+1, divisor)
		}
	}
	return sum
}
