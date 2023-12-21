func GetPermutations(nums *[]int, current_prefix *[]int, permutations *[][]int) {
	if len(*current_prefix) == len(*nums) {
		permutation := make([]int, 0)
		for i := 0; i < len(*current_prefix); i++ {
			permutation = append(permutation, (*current_prefix)[i])
		}
		*permutations = append(*permutations, permutation)
	} else {
		for i := 0; i < len(*nums); i++ {
			if (*nums)[i] != 100 {
				aux := (*nums)[i]
				*current_prefix = append(*current_prefix, aux)
				(*nums)[i] = 100
				GetPermutations(nums, current_prefix, permutations)
				*current_prefix = (*current_prefix)[0 : len(*current_prefix)-1]
				(*nums)[i] = aux
			}
		}
	}
}

func permute(nums []int) [][]int {
	permutations := make([][]int, 0)
	current_prefix := make([]int, 0)
	GetPermutations(&nums, &current_prefix, &permutations)
	return permutations
}
