func lemonadeChange(bills []int) bool {
	cashierBox := map[int]int{}
	for _, bill := range bills {
		if bill == 5 {
			cashierBox[5]++
		} else if bill == 10 {
			if cashierBox[5] == 0 {
				return false
			}
			cashierBox[5]--
			cashierBox[10]++
		} else {
			if 0 < cashierBox[10] && 0 < cashierBox[5] {
				cashierBox[10]--
				cashierBox[5]--
			} else if 3 <= cashierBox[5] {
				cashierBox[5] -= 3
			} else {
				return false
			}
		}
	}
	return true
}
