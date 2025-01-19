func destCity(paths [][]string) string {
	outbounds := map[string]int{}
	for _, path := range paths {
		outbound, inbound := path[0], path[1]
		outbounds[outbound] += 1
		outbounds[inbound] += 0
	}
	for city, count := range outbounds {
		if count == 0 {
			return city
		}
	}
	return ""
}
