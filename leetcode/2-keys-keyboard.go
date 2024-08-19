func minStepsQueue(n int) int {
	visited := map[int]bool{}
	queue := [][]int{{0, 1, 0}} //steps, len, copy
	for 0 < len(queue) {
		entry := queue[0]
		s, l, c := entry[0], entry[1], entry[2]
		key := l*1000 + c
		queue = queue[1:]
		if l == n {
			return s
		}
		if n < l || visited[key] {
			continue
		}
		visited[key] = true
		s++
		queue = append(queue, []int{s, l, l})
		queue = append(queue, []int{s, l + c, c})
	}
	return -1
}

func minStepsRecursive(n, s, l, c int, visited *map[int]bool) int {
	key := l*1000 + c
	if n == l {
		return s
	} else if n < l || (*visited)[key] {
		return 10001
	}
	(*visited)[key] = true
	s++
	return min(minStepsRecursive(n, s, l, l, visited), minStepsRecursive(n, s, l+c, c, visited))
}

func minSteps(n int) int {
	visited := map[int]bool{}
	return minStepsRecursive(n, 0, 1, 0, &visited)
}
