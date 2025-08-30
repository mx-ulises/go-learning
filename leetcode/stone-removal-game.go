package leetcode

func canAliceWin(n int) bool {
	aliceWin, remove := false, 10
	for remove <= n {
		n -= remove
		remove--
		aliceWin = !aliceWin
	}
	return aliceWin
}
