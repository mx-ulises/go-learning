func countSeniors(details []string) int {
	seniors := 0
	for _, passenger := range details {
		if passenger[11] > '6' || (passenger[11] == '6' && passenger[12] > '0') {
			seniors++
		}
	}
	return seniors
}
