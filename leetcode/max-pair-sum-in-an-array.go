package leetcode

func getMaximalDigit(num int) int {
	d := 0
	for 0 < num {
		d = max(d, num%10)
		num /= 10
	}
	return d
}

func maxSum(nums []int) int {
	digitMaximals := [10][2]int{}
	maximal := -1
	for _, num := range nums {
		d := getMaximalDigit(num)
		if digitMaximals[d][1] < num {
			digitMaximal := digitMaximals[d][0]
			digitMaximals[d][0] = max(digitMaximal, num)
			digitMaximals[d][1] = min(digitMaximal, num)
			if digitMaximals[d][1] != 0 {
				maximal = max(maximal, digitMaximal+num)
			}
		}
	}
	return maximal
}
