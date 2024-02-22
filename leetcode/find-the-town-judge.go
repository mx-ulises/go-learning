func findJudge(n int, trust [][]int) int {
	trustCount := map[int]int{}
	trustSomeone := map[int]bool{}
	for _, pair := range trust {
		trustCount[pair[1]]++
		trustSomeone[pair[0]] = true
	}
	judge := -1
	for i := 1; i <= n; i++ {
		if trustCount[i] == (n-1) && trustSomeone[i] == false {
			if judge != -1 {
				return -1
			}
			judge = i
		}
	}
	return judge
}
