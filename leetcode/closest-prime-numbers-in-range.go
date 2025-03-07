func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	i := 2
	for (i * i) <= num {
		if num%i == 0 {
			return false
		}
		i++
	}
	return true
}

func closestPrimes(left int, right int) []int {
	num := left
	a, b := -1, -1
	num1, num2 := -1, -1
	minimal := 1_000_000
	for num <= right {
		if isPrime(num) {
			a, b = b, num
			diff := b - a
			if a != -1 && diff < minimal {
				minimal = diff
				num1, num2 = a, b
				if diff <= 2 {
					break
				}
			}
		}
		num++
	}
	return []int{num1, num2}
}
