type JumbledNumber struct {
	MappedValue int
	Value       int
	Index       int
}

type JumbledNumbers []JumbledNumber

func (a JumbledNumbers) Len() int      { return len(a) }
func (a JumbledNumbers) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a JumbledNumbers) Less(i, j int) bool {
	return a[i].MappedValue < a[j].MappedValue || (a[i].MappedValue == a[j].MappedValue && a[i].Index < a[j].Index)
}

func MapValue(num int, mapping *[]int) int {
	if num == 0 {
		return (*mapping)[0]
	}
	power := 1
	mappedValue := 0
	for 0 < num {
		i := num % 10
		mappedValue += (*mapping)[i] * power
		power *= 10
		num /= 10
	}
	return mappedValue
}

func sortJumbled(mapping []int, nums []int) []int {
	jumbledNumbers := []JumbledNumber{}
	for i, num := range nums {
		jumbledNumbers = append(jumbledNumbers, JumbledNumber{Value: num, Index: i, MappedValue: MapValue(num, &mapping)})
	}
	sort.Sort(JumbledNumbers(jumbledNumbers))
	output := []int{}
	for _, jumbledNumber := range jumbledNumbers {
		output = append(output, jumbledNumber.Value)
	}
	return output
}
