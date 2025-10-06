package leetcode

func recoverOrder(order []int, friends []int) []int {
	output, isFriend, i := make([]int, len(friends)), make([]bool, len(order)+1), 0
	for _, friend := range friends {
		isFriend[friend] = true
	}
	for _, candidate := range order {
		if isFriend[candidate] {
			output[i] = candidate
			i++
		}
	}
	return output
}
