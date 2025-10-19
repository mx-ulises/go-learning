package leetcode

func buyChoco(prices []int, money int) int {
	choco1, choco2 := 101, 101
	for _, choco := range prices {
		if choco <= choco1 {
			choco1, choco2 = choco, choco1
		} else if choco < choco2 {
			choco2 = choco
		}
	}
	leftover := money - choco1 - choco2
	if leftover < 0 {
		leftover = money
	}
	return leftover
}
