package __11

import "github.com/bill_lhr/personal/common"

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	m, n := len(grid), len(grid[0])
	var dfs func(int, int) int

	dfs = func(i int, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0
		return dfs(i-1, j) + dfs(i+1, j) + dfs(i, j-1) + dfs(i, j+1)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				maxArea = common.Max(maxArea, dfs(i, j))
			}
		}
	}

	return maxArea
}
