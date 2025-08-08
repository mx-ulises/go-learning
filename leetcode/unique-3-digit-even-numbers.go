package leetcode

func totalNumbers(digits []int) int {
	digitCount := map[int]int{}
	total := 0
	for _, digit := range digits {
		digitCount[digit]++
	}
	for i := 1; i < 10; i++ {
		if digitCount[i] == 0 {
			continue
		}
		digitCount[i]--
		for j := 0; j < 10; j++ {
			if digitCount[j] == 0 {
				continue
			}
			digitCount[j]--
			for k := 0; k < 10; k += 2 {
				if 0 < digitCount[k] {
					total++
				}
			}
			digitCount[j]++
		}
		digitCount[i]++
	}
	return total
}
