func isValid(i, target int) bool {
	if i == target {
		return true
	}
	p := 1
	for 0 < target {
		i -= p * (target % 10)
		p *= 10
		target /= 10
		if isValid(i, target) {
			return true
		}
	}
	return false
}

func punishmentNumber(n int) int {
	valid := [][]int{
		{1, 1},
		{9, 81},
		{10, 100},
		{36, 1296},
		{45, 2025},
		{55, 3025},
		{82, 6724},
		{91, 8281},
		{99, 9801},
		{100, 10000},
		{235, 55225},
		{297, 88209},
		{369, 136161},
		{370, 136900},
		{379, 143641},
		{414, 171396},
		{657, 431649},
		{675, 455625},
		{703, 494209},
		{756, 571536},
		{792, 627264},
		{909, 826281},
		{918, 842724},
		{945, 893025},
		{964, 929296},
		{990, 980100},
		{991, 982081},
		{999, 998001},
		{1000, 1000000},
	}
	i := 0
	output := 0
	for i < len(valid) && valid[i][0] <= n {
		output += valid[i][1]
		i++
	}
	return output
}
