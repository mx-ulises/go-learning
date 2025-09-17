package leetcode

func lastDigit(num int) int {
	for 10 <= num {
		num /= 10
	}
	return num
}

func gdc(x, y int) int {
	for i := x; 0 < i; i-- {
		if x%i == 0 && y%i == 0 {
			return i
		}
	}
	return -1
}

func getCoprimeMap() [][]bool {
	coprimeMap := make([][]bool, 10)
	for x := 0; x < 10; x++ {
		coprimeMap[x] = make([]bool, 10)
		for y := x; y < 10; y++ {
			if gdc(x, y) == 1 {
				coprimeMap[x][y] = true
			}
		}
	}
	return coprimeMap
}

func countBeautifulPairs(nums []int) int {
	count, n := 0, len(nums)
	coprimeMap := getCoprimeMap()
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a, b := lastDigit(nums[i]), nums[j]%10
			x, y := min(a, b), max(a, b)
			if coprimeMap[x][y] {
				count++
			}
		}
	}
	return count
}
