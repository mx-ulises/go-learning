func isBalanced(num string) bool {
	sumMap := [2]rune{}
	i := 0
	for _, d := range num {
		sumMap[i] += (d - '0')
		i ^= 1
	}
	fmt.Println(sumMap)
	return sumMap[0] == sumMap[1]
}
