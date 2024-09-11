func minBitFlips(start int, goal int) int {
	diff := start ^ goal
	flips := 0
	for diff != 0 {
		flips += (diff & 1)
		diff >>= 1
	}
	return flips
}
