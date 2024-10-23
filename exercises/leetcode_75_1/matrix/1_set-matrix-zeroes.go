/*
	Problem: 73. Set Matrix Zeroes
	Reference: https://leetcode.com/problems/set-matrix-zeroes/description/
    Complexity: O(m*n)
*/

func setZeroes(matrix [][]int) {
	row := make(map[int]struct{}, len(matrix))
	col := make(map[int]struct{}, len(matrix[0]))
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				row[i] = struct{}{}
				col[j] = struct{}{}
			}
		}
	}
	for i := range matrix {
		for j := range matrix[i] {
			_, okRow := row[i]
			_, okCol := col[j]
			if okRow || okCol {
				matrix[i][j] = 0
			}
		}
	}
}