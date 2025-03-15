func getLaserCount(row string) int {
	laserCount := 0
	for _, c := range row {
		if c == '1' {
			laserCount++
		}
	}
	return laserCount
}

func numberOfBeams(bank []string) int {
	prevLaserCount := 0
	count := 0
	for _, row := range bank {
		currentLaserCount := getLaserCount(row)
		if currentLaserCount != 0 {
			count += prevLaserCount * currentLaserCount
			prevLaserCount = currentLaserCount
		}
	}
	return count
}
