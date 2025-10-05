package leetcode

func getNum(date string) int {
	return 10*int(date[0]-'0') + int(date[1]-'0')
}

func getDayNumber(calendar []int, date string) int {
	month, day := getNum(date[0:2])-1, getNum(date[3:5])
	return calendar[month] + day
}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	calendar := []int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
	start1, end1 := getDayNumber(calendar, arriveAlice), getDayNumber(calendar, leaveAlice)
	start2, end2 := getDayNumber(calendar, arriveBob), getDayNumber(calendar, leaveBob)
	return max(0, min(end1, end2)-max(start1, start2)+1)
}
