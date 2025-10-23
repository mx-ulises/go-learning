package leetcode

func getMinutes(ts string) int {
	h := 10*int(byte(ts[0]-'0')) + int(byte(ts[1]-'0'))
	m := 10*int(byte(ts[3]-'0')) + int(byte(ts[4]-'0'))
	return 60*h + m
}

func getDivAndModulo(ops, diff, unit int) (int, int) {
	return ops + diff/unit, diff % unit
}

func convertTime(current string, correct string) int {
	ops, diff := 0, getMinutes(correct)-getMinutes(current)
	ops, diff = getDivAndModulo(ops, diff, 60)
	ops, diff = getDivAndModulo(ops, diff, 15)
	ops, diff = getDivAndModulo(ops, diff, 5)
	return ops + diff
}
