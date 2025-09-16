package leetcode

func earliestTime(tasks [][]int) int {
	earliest := tasks[0][0] + tasks[0][1]
	for _, task := range tasks {
		earliest = min(earliest, task[0]+task[1])
	}
	return earliest
}
