package leetcode

func checkString(s string) bool {
	bSeen := false
	for _, c := range s {
		if c == 'b' {
			bSeen = true
		} else if bSeen {
			return false
		}
	}
	return true
}
