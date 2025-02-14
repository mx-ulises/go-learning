func deleteGreatestValue(grid [][]int) int {
    n, m := len(grid), len(grid[0])
    answer := 0
    for i := 0; i < n; i++ {
        sort.Ints(grid[i])
    }
    for j := 0; j < m; j++ {
        maximal := 0
        for i := 0; i < n; i++ {
            maximal = max(grid[i][m-j-1], maximal)
        }
        answer += maximal
    }
    return answer
}
