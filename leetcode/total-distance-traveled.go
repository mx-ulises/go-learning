package leetcode

func algoSolution(mainTank int, additionalTank int) int {
	totalLitres := 0
	for 5 <= mainTank {
		remaining := mainTank % 5
		transferred := min(mainTank/5, additionalTank)
		totalLitres += mainTank - remaining
		additionalTank -= transferred
		mainTank = remaining + transferred
	}
	return totalLitres + mainTank
}

func mathSolution(mainTank int, additionalTank int) int {
	return (mainTank + min((mainTank-1)/4, additionalTank))
}

func distanceTraveled(mainTank int, additionalTank int) int {
	return mathSolution(mainTank, additionalTank) * 10
}
