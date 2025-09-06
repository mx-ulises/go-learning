package leetcode

func distMoney(money int, children int) int {
	money -= children
	count := min(money/7, children)
	remaining := money - count*7
	if money < 0 {
		count = -1
	} else if count == children && 0 < remaining {
		count--
	} else if count == (children-1) && 3 == remaining {
		count--
	}
	return count
}
