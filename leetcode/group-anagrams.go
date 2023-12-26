func SortString(s string) string {
	characters := strings.Split(s, "")
	sort.Strings(characters)
	return strings.Join(characters, "")
}

func groupAnagrams(strs []string) [][]string {
	stringGroups := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		s := strs[i]
		key := SortString(s)
		stringGroups[key] = append(stringGroups[key], s)
	}
	output := make([][]string, 0)
	for _, stringGroup := range stringGroups {
		output = append(output, stringGroup)
	}
	return output
}
