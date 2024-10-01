func canArrange(arr []int, k int) bool {
	modCount := map[int]int{}
	for _, num := range arr {
		if num < 0 {
			times := (num / -k) + 1
			num += k * times
		}
		modCount[num%k]++
	}
	m := k / 2
	for i := 1; i <= m; i++ {
		if modCount[i] != modCount[k-i] {
			return false
		}
	}
	if k&1 == 0 && modCount[m]&1 != 0 {
		return false
	}
	return modCount[0]&1 == 0
}
