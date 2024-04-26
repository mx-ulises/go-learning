func GetMultiplerSum(increment int, n int, exceptions []int) int {
	sum, i := 0, 0
	for i <= n {
		isException := false
		for _, j := range exceptions {
			if i%j == 0 {
				isException = true
				break
			}
		}
		if isException == false {
			sum += i
		}
		i += increment
	}
	return sum
}

func sumOfMultiples(n int) int {
	exceptions := []int{}
	sum := GetMultiplerSum(3, n, exceptions)
	exceptions = append(exceptions, 3)
	sum += GetMultiplerSum(5, n, exceptions)
	exceptions = append(exceptions, 5)
	sum += GetMultiplerSum(7, n, exceptions)
	return sum
}
