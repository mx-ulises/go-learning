package leetcode

func numPrimeArrangements(n int) int {
	primeCount := getPrimeCountUpTo(n)
	nonPrimeCount := n - primeCount
	result := getPermutations(primeCount) * getPermutations(nonPrimeCount)
	return result % 1_000_000_007
}

func getPrimeCountUpTo(n int) int {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	for i := 0; i < len(primes); i++ {
		if n < primes[i] {
			return i
		}
	}
	return len(primes)
}

func getPermutations(n int) int {
	return getFactorial(n)
}

func getFactorial(n int) int {
	factorial := 1
	for i := 2; i <= n; i++ {
		factorial *= i
		factorial %= 1_000_000_007
	}
	return factorial
}
