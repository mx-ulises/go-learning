package leetcode

func slowestKey(releaseTimes []int, keysPressed string) byte {
	n, key, slowest, prev := len(releaseTimes), keysPressed[0], releaseTimes[0], releaseTimes[0]
	for i := 1; i < n; i++ {
		candidate := releaseTimes[i] - prev
		if slowest < candidate {
			slowest, key = candidate, keysPressed[i]
		} else if slowest == candidate {
			key = max(key, keysPressed[i])
		}
		prev = releaseTimes[i]
	}
	return key
}
