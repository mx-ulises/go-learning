func stackSolution(arr []int) int {
	currentSum := arr[0]
	totalSum := 0
	s := [][2]int{{arr[0], 1}}
	n := 0
	for i := 1; i < len(arr); i++ {
		totalSum += currentSum
		count := 1
		for 0 <= n && arr[i] < s[n][0] {
			element := s[n]
			s = s[:n]
			count += element[1]
			currentSum -= element[0] * element[1]
			n--
		}
		s = append(s, [2]int{arr[i], count})
		n++
		currentSum += arr[i] * count
	}
	totalSum += currentSum
	return totalSum % 1_000_000_007
}

func bruteForce(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		minimal := arr[i]
		for j := i; j < len(arr); j++ {
			if arr[j] < minimal {
				minimal = arr[j]
			}
			sum += minimal
		}
	}
	return sum % 1_000_000_007
}

func sumSubarrayMins(arr []int) int {
	//return bruteForce(arr)
	return stackSolution(arr)
}
