func sum(skill []int) int {
	total := 0
	for _, member := range skill {
		total += member
	}
	return total
}

func dividePlayers(skill []int) int64 {
	m := len(skill) / 2
	teamSkill := sum(skill) / m
	chemestry := 0
	members := map[int]int{}
	for _, member := range skill {
		memberMatch := teamSkill - member
		if 0 < members[memberMatch] {
			members[memberMatch]--
			m--
			chemestry += member * memberMatch
		} else {
			members[member]++
		}
	}
	if m != 0 {
		return -1
	}
	return int64(chemestry)
}
