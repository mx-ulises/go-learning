func removePalindromeSub(s string) int {
	deleted, n, removals := 0, len(s), 0
	visited := make([]bool, n)
	chars := map[byte]bool{}
	for deleted < n {
		removals++
		left, right := 0, n-1
		for left <= right {
			for left <= right && visited[left] {
				left++
			}
			if !visited[left] {
				chars[s[left]] = true
			}
			for left <= right && visited[right] {
				right--
			}
			if s[left] == s[right] {
				deleted += 2
				if left == right {
					deleted--
				}
				visited[left], visited[right] = true, true
				left, right = left+1, right-1
			} else {
				left++
			}
		}
	}
	return min(removals, len(chars))
}
