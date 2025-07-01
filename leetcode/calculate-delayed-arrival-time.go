package leetcode

func findDelayedArrivalTime(arrivalTime int, delayedTime int) int {
	arrivalTime += delayedTime
	if 24 <= arrivalTime {
		arrivalTime -= 24
	}
	return arrivalTime
}
