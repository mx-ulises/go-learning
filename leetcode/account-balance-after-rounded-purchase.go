package leetcode

func accountBalanceAfterPurchase(purchaseAmount int) int {
	d := purchaseAmount % 10
	purchaseAmount -= d
	if d < 5 {
		d = 0
	} else {
		d = 10
	}
	return 100 - (purchaseAmount + d)
}
