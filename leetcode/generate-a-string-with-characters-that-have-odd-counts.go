func generateTheString(n int) string {
	output := make([]byte, n)
	for i := 0; i < n; i++ {
		output[i] = 'a'
	}
	if n&1 == 0 {
		output[0] = 'b'
	}
	return string(output)

}
