package leetcode

func checkOnesSegment(s string) bool {
	segmentCount := initializeSegmentCount(s[0])
	for i := 1; i < len(s); i++ {
		segmentCount = updateSegmentCount(s[i], s[i-1], segmentCount)
	}
	return segmentCount <= 1
}

func initializeSegmentCount(firstBit byte) int {
	if firstBit == '1' {
		return 1
	}
	return 0
}

func updateSegmentCount(currentBit byte, previousBit byte, segmentCount int) int {
	if isStartingNewSegment(currentBit, previousBit) {
		segmentCount++
	}
	return segmentCount
}

func isStartingNewSegment(currentBit byte, previousBit byte) bool {
	return currentBit == '1' && previousBit == '0'
}
