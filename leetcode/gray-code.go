func GetGrayCode(n int, visited *map[int]bool, output *[]int, x int) {
	if (*visited)[x] == false {
		(*visited)[x] = true
		(*output) = append((*output), x)
		mask := 1
		for i := 0; i < n; i++ {
			GetGrayCode(n, visited, output, x^mask)
			mask = mask << 1
		}
	}
}

func grayCodeRecursive(n int) []int {
	visited := make(map[int]bool)
	output := make([]int, 0)
	GetGrayCode(n, &visited, &output, 0)
	return output
}

func grayCodeEfficient(n int) []int {
	codes := make([]int, 2<<(n-1))
	for i := 0; i < len(codes); i++ {
		codes[i] = i ^ (i >> 1)
	}
	return codes
}

func grayCodeIterative(n int) []int {
	visited := make(map[int]bool)
	output := make([]int, 0)
	s := []int{0}
	for 0 < len(s) {
		l := len(s) - 1
		x := s[l]
		s = s[:l]
		if visited[x] {
			continue
		}
		visited[x] = true
		output = append(output, x)
		mask := 1
		for i := 0; i < n; i++ {
			s = append(s, x^mask)
			mask = mask << 1
		}
	}
	return output
}

func grayCode(n int) []int {
	return grayCodeEfficient(n)
}
