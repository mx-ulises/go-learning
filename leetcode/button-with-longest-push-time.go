package leetcode

func buttonWithLongestTime(events [][]int) int {
	p, wait := 0, -1
	button := -1
	longestWait := 0
	longestWaitPerButton := map[int]int{}
	for _, event := range events {
		i, t := event[0], event[1]
		wait = t - p
		p = t
		longestWaitPerButton[i] = max(longestWaitPerButton[i], wait)
		if longestWait < longestWaitPerButton[i] {
			longestWait = longestWaitPerButton[i]
			button = i
		}
		if longestWait == longestWaitPerButton[i] {
			button = min(button, i)
		}
	}
	return button
}
