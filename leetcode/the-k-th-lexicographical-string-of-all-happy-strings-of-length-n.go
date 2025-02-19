func getTotalStrings(n int) int {
	suffix := math.Pow(2.0, float64(n-1))
	totalStrings := 3 * int(suffix)
	return totalStrings
}

func getBucketAndK(totalStrings, k, optionCount int) (int, int) {
	bucketSize := totalStrings / optionCount
	bucket := (k - 1) / bucketSize
	k -= bucket * bucketSize
	return bucket, k
}

func getHappyString(n int, k int) string {
	totalStrings := getTotalStrings(n)
	if totalStrings < k {
		return ""
	}
	OPTIONS := map[byte][]byte{
		' ': {'a', 'b', 'c'},
		'a': {'b', 'c'},
		'b': {'a', 'c'},
		'c': {'a', 'b'},
	}
	var prev byte = ' '
	s := []byte{}
	for len(s) < n {
		var bucket int
		bucket, k = getBucketAndK(totalStrings, k, len(OPTIONS[prev]))
		s = append(s, OPTIONS[prev][bucket])
		totalStrings /= len(OPTIONS[prev])
		prev = OPTIONS[prev][bucket]
	}
	return string(s)
}
