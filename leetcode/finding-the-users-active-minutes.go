func findingUsersActiveMinutes(logs [][]int, k int) []int {
	userMinutes := map[int]map[int]bool{}
	for _, log := range logs {
		uid, time := log[0], log[1]
		_, ok := userMinutes[uid]
		if !ok {
			userMinutes[uid] = map[int]bool{}
		}
		userMinutes[uid][time] = true
	}
	userActiveMinutes := make([]int, k)
	for _, times := range userMinutes {
		activeMinutes := len(times) - 1
		userActiveMinutes[activeMinutes] += 1
	}
	return userActiveMinutes
}
