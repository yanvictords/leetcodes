/*
	Problem: 200. Number of Islands
	Reference: https://leetcode.com/problems/number-of-islands
    Complexity: O(V+E)
*/

func eraseIsland(grid [][]byte, i, j int) {
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == '0' {
        return
    }

    grid[i][j] = '0'
    eraseIsland(grid, i,j-1)
    eraseIsland(grid, i,j+1)
    eraseIsland(grid, i-1,j)
    eraseIsland(grid, i+1,j)
}

func numIslands(grid [][]byte) int {
    total := 0
    for i, row := range grid {
        for j, node := range row {
            if node == '1' {
                total++
                eraseIsland(grid, i, j)
            }
        }
    }
    return total
}