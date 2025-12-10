package leetcode

import "sort"

func sortByReflection(nums []int) []int {
	binaryReflectionMap := getBinaryReflectionMap(nums)
	return getSortedArray(binaryReflectionMap)
}

func getBinaryReflectionMap(nums []int) map[int][]int {
	binaryReflectionMap := map[int][]int{}
	for _, num := range nums {
		binaryReflectionMap = appendInBinaryReflectionMap(binaryReflectionMap, num)
	}
	return binaryReflectionMap
}

func appendInBinaryReflectionMap(binaryReflectionMap map[int][]int, num int) map[int][]int {
	binaryReflection := getBinaryReflection(num)
	binaryReflectionMap[binaryReflection] = append(binaryReflectionMap[binaryReflection], num)
	return binaryReflectionMap
}

func getBinaryReflection(num int) int {
	binaryReflection := 0
	for 0 < num {
		binaryReflection = appendBit(binaryReflection, num&1)
		num >>= 1
	}
	return binaryReflection
}

func appendBit(binaryReflection int, b int) int {
	return (binaryReflection << 1) + b
}

func getSortedArray(binaryReflectionMap map[int][]int) []int {
	sortedBinaryReflections := getSortedBinaryReflections(binaryReflectionMap)
	return mergeSortedBuckets(binaryReflectionMap, sortedBinaryReflections)
}

func getSortedBinaryReflections(binaryReflectionMap map[int][]int) []int {
	binaryReflections := getBinaryReflections(binaryReflectionMap)
	sort.Ints(binaryReflections)
	return binaryReflections
}

func mergeSortedBuckets(binaryReflectionMap map[int][]int, sortedBinaryReflections []int) []int {
	output := []int{}
	for _, binaryReflection := range sortedBinaryReflections {
		output = sortAndAppendBucket(output, binaryReflectionMap[binaryReflection])
	}
	return output
}

func sortAndAppendBucket(output []int, bucket []int) []int {
	sort.Ints(bucket)
	return append(output, bucket...)
}

func getBinaryReflections(binaryReflectionMap map[int][]int) []int {
	binaryReflections := []int{}
	for binaryReflection, _ := range binaryReflectionMap {
		binaryReflections = append(binaryReflections, binaryReflection)
	}
	return binaryReflections
}
