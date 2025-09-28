package leetcode

func dayOfTheWeek(day int, month int, year int) string {
	weekdays := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	months := []int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
	startDay := 4
	startDay += 365*(year-1971) + (year-1968-1)/4 // Days of the year
	startDay += months[month-1]                   // Days of the month
	startDay += day                               // current day
	if year%4 == 0 && 2 < month && year != 2100 {
		startDay++ // leap year extra day
	}
	return weekdays[startDay%7]
}
