func maximum69Number(num int) int {
	aux, maximal, add := num, num, 3
	for 0 < aux {
		if aux%10 == 6 {
			maximal = num + add
		}
		add *= 10
		aux /= 10
	}
	return maximal
}
