package leetcode

func parity(c byte, offset byte) byte {
	return (c - offset) & 1
}

func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
	parityMix1 := parity(coordinate1[0], 'a') ^ parity(coordinate1[1], '1')
	parityMix2 := parity(coordinate2[0], 'a') ^ parity(coordinate2[1], '1')
	return parityMix1 == parityMix2
}
