type Jump struct {
	Position int
	Steps    int
}

func jumpBFS(nums []int) int {
	queue := make([]Jump, 0)
	visited := make([]bool, len(nums))
	target := len(nums) - 1
	queue = append(queue, Jump{Position: 0, Steps: 0})
	for 0 < len(queue) {
		jump := queue[0]
		queue = queue[1:]
		if jump.Position == target {
			return jump.Steps
		}
		if target < jump.Position || visited[jump.Position] {
			continue
		}
		visited[jump.Position] = true
		for i := nums[jump.Position]; 0 < i; i-- {
			queue = append(queue, Jump{Position: jump.Position + i, Steps: jump.Steps + 1})
		}
	}
	return -1
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func jump(nums []int) int {
	jumpCount := make([]int, len(nums))
	for i := 1; i < len(jumpCount); i++ {
		jumpCount[i] = 10000
	}
	maxJumpPosition := make([]int, 0)
	maxJumpPosition = append(maxJumpPosition, 0)
	for i := 0; i < len(nums); i++ {
		jumps := jumpCount[i] + 1
		if len(maxJumpPosition) == jumps {
			maxJumpPosition = append(maxJumpPosition, i+1)
		}
		limit := min(i+nums[i], len(nums)-1)
		for position := maxJumpPosition[jumps]; position <= limit; position++ {
			jumpCount[position] = min(jumpCount[position], jumps)
		}
		maxJumpPosition[jumps] = max(limit, maxJumpPosition[jumps])
	}
	return jumpCount[len(nums)-1]
}
