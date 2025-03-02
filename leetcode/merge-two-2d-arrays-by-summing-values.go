func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	i, j := 0, 0
	n, m := len(nums1), len(nums2)
	output := [][]int{}
	for i < n || j < m {
		if i == n {
			output = append(output, nums2[j])
			j++
		} else if j == m {
			output = append(output, nums1[i])
			i++
		} else if nums1[i][0] < nums2[j][0] {
			output = append(output, nums1[i])
			i++
		} else if nums2[j][0] < nums1[i][0] {
			output = append(output, nums2[j])
			j++
		} else {
			output = append(output, nums1[i])
			nums1[i][1] += nums2[j][1]
			i++
			j++
		}
	}
	return output
}
