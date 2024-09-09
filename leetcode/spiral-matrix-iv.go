/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func GetMatrix(m, n int) *[][]int {
	output := make([][]int, m)
	for i := 0; i < m; i++ {
		output[i] = make([]int, n)
		for j := 0; j < n; j++ {
			output[i][j] = -1
		}
	}
	return &output
}

func IsValid(x int, y int, output *[][]int) bool {
	if x < 0 || x == len(*output) {
		return false
	}
	if y < 0 || y == len((*output)[0]) {
		return false
	}
	return (*output)[x][y] == -1
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	output := GetMatrix(m, n)
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	x, y, d := 0, 0, 0
	for head != nil {
		(*output)[x][y] = head.Val
		if IsValid(x+directions[d][0], y+directions[d][1], output) == false {
			d = (d + 1) % 4
		}
		x += directions[d][0]
		y += directions[d][1]
		head = head.Next
	}
	return *output
}
