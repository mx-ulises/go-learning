func merge(nums1 []int, m int, nums2 []int, n int) {
	l := len(nums1) - 1
	m--
	n--
	for 0 <= l {
		if m < 0 {
			nums1[l] = nums2[n]
			n--
		} else if n < 0 {
			nums1[l] = nums1[m]
			m--
		} else if nums1[m] < nums2[n] {
			nums1[l] = nums2[n]
			n--
		} else {
			nums1[l] = nums1[m]
			m--
		}
		l--
	}
}
