func judgeSquareSum2(c int) bool {
	candidates := []int{}
	x, p := 0, 0
	for p <= c {
		candidates = append(candidates, p)
		x++
		p = x * x
	}
	i, j := 0, x-1
	for i <= j {
		total := candidates[i] + candidates[j]
		if total == c {
			return true
		} else if total < c {
			i += 1
		} else {
			j -= 1
		}
	}
	return false
}

func judgeSquareSum(c int) bool {
	visited := map[int]bool{}
	x, p := 0, 0
	for p <= c {
		p = x * x
		visited[p] = true
		if visited[c-p] {
			return true
		}
		x++
	}
	return false
}
