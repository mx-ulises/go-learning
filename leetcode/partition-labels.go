type ByFirstElement [][]int

func (a ByFirstElement) Len() int           { return len(a) }
func (a ByFirstElement) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFirstElement) Less(i, j int) bool { return a[i][0] < a[j][0] }

func getRanges(s string) [][]int {
	rangesMap := map[rune][]int{}
	for i, c := range s {
		_, ok := rangesMap[c]
		if ok == false {
			rangesMap[c] = []int{i, i}
		}
		rangesMap[c][1] = i
	}
	ranges := make([][]int, len(rangesMap))
	i := 0
	for _, r := range rangesMap {
		ranges[i] = r
		i++
	}
	sort.Sort(ByFirstElement(ranges))
	return ranges
}

func partitionLabels(s string) []int {
	ranges := getRanges(s)
	output := []int{}
	start, end := ranges[0][0], ranges[0][1]
	for i := 1; i < len(ranges); i++ {
		if ranges[i][0] <= end {
			end = max(end, ranges[i][1])
		} else {
			output = append(output, end-start+1)
			start, end = ranges[i][0], ranges[i][1]
		}
	}
	output = append(output, end-start+1)
	return output
}
