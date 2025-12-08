package leetcode

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	earliestFinishTimeStartingWithLand := getEarliestFinishTimeStartingWithLand(landStartTime, landDuration, waterStartTime, waterDuration)
	earliestFinishTimeStartingWithWater := getEarliestFinishTimeStartingWithWater(landStartTime, landDuration, waterStartTime, waterDuration)
	return min(earliestFinishTimeStartingWithLand, earliestFinishTimeStartingWithWater)
}

func getEarliestFinishTimeStartingWithLand(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	return getEarliestFinishTime(landStartTime, landDuration, waterStartTime, waterDuration)
}

func getEarliestFinishTimeStartingWithWater(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	return getEarliestFinishTime(waterStartTime, waterDuration, landStartTime, landDuration)
}

func getEarliestFinishTime(firstRideStartTime []int, firstRideDuration []int, secondRideStartTime []int, secondRideDuration []int) int {
	firstRideFinishTime := getEarliestRideFinishTime(firstRideStartTime, firstRideDuration, 0)
	secondRideFinishTime := getEarliestRideFinishTime(secondRideStartTime, secondRideDuration, firstRideFinishTime)
	return secondRideFinishTime
}

func getEarliestRideFinishTime(rideStartTime []int, rideDuration []int, earliestStartTime int) int {
	earliestRideFinishTime := 3000
	for i := 0; i < len(rideStartTime); i++ {
		rideFinishTime := max(earliestStartTime, rideStartTime[i]) + rideDuration[i]
		earliestRideFinishTime = min(earliestRideFinishTime, rideFinishTime)
	}
	return earliestRideFinishTime
}
