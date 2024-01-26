func minPartitions(n string) int {
	partitions := 0
	for _, c := range n {
		d := int(c - '0')
		if partitions < d {
			partitions = d
		}
	}
	return partitions
}
