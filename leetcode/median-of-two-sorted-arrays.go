func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Stat index (local and global), limits, midpoints and output values
	i1, i2, n1, n2 := 0, 0, len(nums1), len(nums2)
	j, n := 0, n1+n2
	m1, m2 := (n-1)/2, (n)/2
	var output1, output2 float64
	// Continue until break
	for true {
		// Get current num and update local index
		var num float64
		if i1 == n1 || (i2 < n2 && nums2[i2] < nums1[i1]) {
			num = float64(nums2[i2])
			i2++
		} else {
			num = float64(nums1[i1])
			i1++
		}
		// Fill means as needed and break
		if j == m1 {
			output1 = num
		}
		if j == m2 {
			output2 = num
			break
		}
		// Update global index
		j = i1 + i2
	}
	return (output1 + output2) / 2
}
