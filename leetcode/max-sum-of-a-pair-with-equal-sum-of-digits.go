func getSumKey(num int) int {
    sumKey := 0
    for 0 < num {
        sumKey += num % 10
        num /= 10
    }
    return sumKey
}

func maximumSum(nums []int) int {
    sumToMax := map[int]int{}
    maximal := -1
    for _, num := range nums {
        sumKey := getSumKey(num)
        maximalWithSum, ok := sumToMax[sumKey]
        if ok == true {
            maximal = max(maximal, maximalWithSum + num)
        }
        sumToMax[sumKey] = max(sumToMax[sumKey], num)
    }
    return maximal
}
