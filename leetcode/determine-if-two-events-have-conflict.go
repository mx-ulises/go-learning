package leetcode

func getMinutes(ts string) int {
	hours := 10*int(ts[0]) + int(ts[1])
	minutes := 10*int(ts[3]) + int(ts[4])
	return 60*hours + minutes
}

func testTime(s, e, m int) bool {
	return s <= m && m <= e
}

func haveConflict(event1 []string, event2 []string) bool {
	s1, e1 := getMinutes(event1[0]), getMinutes(event1[1])
	s2, e2 := getMinutes(event2[0]), getMinutes(event2[1])
	return testTime(s1, e1, s2) || testTime(s2, e2, s1)
}
