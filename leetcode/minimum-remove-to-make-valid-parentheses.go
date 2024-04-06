func GetIndexesToRemove(s string) []int {
	indexes := []int{}
	for i, c := range s {
		n := len(indexes)
		if c != '(' && c != ')' {
			continue
		} else if c == ')' && 0 < n && s[indexes[n-1]] == '(' {
			indexes = indexes[:n-1]
		} else {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func RemoveIndexes(s string, indexes []int) string {
	output := []rune{}
	for i, c := range s {
		if 0 < len(indexes) && indexes[0] == i {
			indexes = indexes[1:]
		} else {
			output = append(output, c)
		}
	}
	return string(output)
}

func minRemoveToMakeValid(s string) string {
	indexes := GetIndexesToRemove(s)
	return RemoveIndexes(s, indexes)
}
