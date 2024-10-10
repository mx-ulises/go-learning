func GetNumbers(s string) (int64, int64, int64) {
	current := []rune{}
	nums := []int64{}
	for _, c := range s {
		if c == '-' {
			num, _ := strconv.Atoi(string(current))
			nums = append(nums, int64(num))
			current = []rune{}
		} else {
			current = append(current, c)
		}
	}
	num, _ := strconv.Atoi(string(current))
	nums = append(nums, int64(num))
	return nums[0], nums[1], nums[2]
}

func convertDateToBinary(date string) string {
	year, month, day := GetNumbers(date)
	return strconv.FormatInt(year, 2) + "-" + strconv.FormatInt(month, 2) + "-" + strconv.FormatInt(day, 2)
}
