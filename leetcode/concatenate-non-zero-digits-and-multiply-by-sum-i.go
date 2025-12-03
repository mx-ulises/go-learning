package leetcode

func sumAndMultiply(n int) int64 {
	concatenation := getNonZeroConcatenation(n)
	sum := getDigitSum(n)
	return concatenation * sum
}

func getNonZeroConcatenation(n int) int64 {
	concatenationReversed := 0
	for 0 < n {
		concatenationReversed = appendNonZero(concatenationReversed, n%10)
		n /= 10
	}
	concatenation := reverseDigits(concatenationReversed)
	return int64(concatenation)
}

func appendNonZero(concatenation int, d int) int {
	if d != 0 {
		concatenation = concatenation*10 + d
	}
	return concatenation
}

func reverseDigits(concatenationReversed int) int {
	concat := 0
	for 0 < concatenationReversed {
		concat = concat*10 + concatenationReversed%10
		concatenationReversed /= 10
	}
	return concat
}

func getDigitSum(n int) int64 {
	sum := 0
	for 0 < n {
		sum += n % 10
		n /= 10
	}
	return int64(sum)
}
