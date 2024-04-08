func GetCountOfStudentThatWantsSandwich(sandwiches []int) [2]int {
	countOfStudentThatWantsSandwich := [2]int{0, 0}
	for _, sandwich := range sandwiches {
		countOfStudentThatWantsSandwich[sandwich]++
	}
	return countOfStudentThatWantsSandwich
}

func countStudents(students []int, sandwiches []int) int {
	countOfStudentThatWantsSandwich := GetCountOfStudentThatWantsSandwich(students)
	i, n := 0, len(sandwiches)
	for i < n && countOfStudentThatWantsSandwich[sandwiches[i]] != 0 {
		countOfStudentThatWantsSandwich[sandwiches[i]]--
		i++
	}
	return countOfStudentThatWantsSandwich[0] + countOfStudentThatWantsSandwich[1]
}
