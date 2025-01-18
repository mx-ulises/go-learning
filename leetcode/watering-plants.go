func wateringPlants(plants []int, capacity int) int {
	steps, bucket := 0, capacity
	for i, plant := range plants {
		if bucket < plant {
			bucket = capacity
			steps += 2 * i
		}
		bucket -= plant
		steps++
	}
	return steps
}
