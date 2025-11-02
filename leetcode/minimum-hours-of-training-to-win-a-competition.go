package leetcode

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	training, n := 0, len(energy)
	for i := 0; i < n; i++ {
		if initialEnergy <= energy[i] {
			training += (energy[i] - initialEnergy) + 1
			initialEnergy = 1
		} else {
			initialEnergy -= energy[i]
		}
		if initialExperience <= experience[i] {
			training += (experience[i] - initialExperience) + 1
			initialExperience = 2*experience[i] + 1
		} else {
			initialExperience += experience[i]
		}
	}
	return training
}
