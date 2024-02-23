func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	employeesOnTarget := 0
	for _, h := range hours {
		if target <= h {
			employeesOnTarget++
		}
	}
	return employeesOnTarget
}
