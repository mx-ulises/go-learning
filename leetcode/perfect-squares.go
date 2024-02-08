type NumSum struct {
	Num int
	Sum int
}

func numSquares(n int) int {
	options := map[int]bool{}
	visited := map[int]bool{}
	queue := make([]NumSum, 0)
	k := 1
	for i := 1; k <= n; i++ {
		k = i * i
		options[k] = true
		queue = append(queue, NumSum{Num: k, Sum: 1})
	}
	for 0 < len(queue) {
		item := queue[0]
		queue = queue[1:]
		if item.Num == n {
			return item.Sum
		}
		if visited[item.Num] {
			continue
		}
		visited[item.Num] = true
		for k, _ := range options {
			j := k + item.Num
			if j <= n {
				queue = append(queue, NumSum{Num: j, Sum: item.Sum + 1})
			}
		}
	}
	return -1
}
