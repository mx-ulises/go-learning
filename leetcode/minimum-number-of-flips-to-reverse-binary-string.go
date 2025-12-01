package leetcode

func minimumFlips(original int) int {
	reversed := getReversed(original)
	diff := original ^ reversed
	return countDifferences(diff)
}

func getReversed(original int) int {
	reversed := 0
	for 0 < original {
		reversed <<= 1
		reversed += original & 1
		original >>= 1
	}
	return reversed
}

func countDifferences(diff int) int {
	requiredFlips := 0
	for 0 < diff {
		requiredFlips += diff & 1
		diff >>= 1
	}
	return requiredFlips
}
