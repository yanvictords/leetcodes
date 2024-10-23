/*
	Problem: 62. Unique Paths
	Reference: https://leetcode.com/problems/unique-paths/
	Complexity: O(M*N)
*/

func dfs(m int, n int, x, y int, grid [][]int) int {
	if x == m || y == n {
		return 0
	}
	if x == m-1 && y == n-1 {
		return 0
	}
	if x == m-2 && y == n-1 {
		return 1
	}
	if x == m-1 && y == n-2 {
		return 1
	}

	if grid[x][y] != 0 {
		return grid[x][y]
	}

	grid[x][y] = dfs(m, n, x+1, y, grid) + dfs(m, n, x, y+1, grid)

	return grid[x][y]
}

func uniquePaths(m int, n int) int {
	if m == 1 && n == 1 {
		return 1
	}
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
	}
	return dfs(m, n, 0, 0, grid)
}