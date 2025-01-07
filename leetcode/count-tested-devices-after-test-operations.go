func countTestedDevices(batteryPercentages []int) int {
	decrease, n := 0, len(batteryPercentages)
	for i := 0; i < n; i++ {
		if decrease < batteryPercentages[i] {
			decrease++
		}
	}
	return decrease
}
