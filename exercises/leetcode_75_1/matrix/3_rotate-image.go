/*
	Problem: 48. Rotate Image
	Reference: https://leetcode.com/problems/rotate-image/description/
    Complexity: O(m*n)
*/

func invert(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i])/2; j++ {
			last := len(matrix[i]) - 1 - j
			matrix[i][j], matrix[i][last] = matrix[i][last], matrix[i][j]
		}
	}
	return matrix
}

func transpose(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

func rotate(matrix [][]int) {
	invert(transpose(matrix))
}