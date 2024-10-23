/*
	Problem: 54. Spiral Matrix
	Reference: https://leetcode.com/problems/spiral-matrix/
    Complexity: O(m*n)
*/

const (
	R       = 1
	D       = 2
	L       = 3
	U       = 4
	VISITED = -100
)

func spiralOrder(matrix [][]int) []int {
	dir := R
	ans := []int{}
	row, col := 0, 0
	for row >= 0 && col >= 0 &&
		row < len(matrix) &&
		col < len(matrix[row]) &&
		matrix[row][col] != VISITED {

		ans = append(ans, matrix[row][col])
		matrix[row][col] = VISITED
		if dir == R {
			if col != len(matrix[row])-1 && matrix[row][col+1] != VISITED {
				col++
				continue
			}
			dir = D
			row++
			continue
		}
		if dir == D {
			if row != len(matrix)-1 && matrix[row+1][col] != VISITED {
				row++
				continue
			}
			dir = L
			col--
			continue
		}
		if dir == L {
			if col != 0 && matrix[row][col-1] != VISITED {
				col--
				continue
			}
			dir = U
			row--
			continue
		}
		if dir == U {
			if row != 0 && matrix[row-1][col] != VISITED {
				row--
				continue
			}
			dir = R
			col++
			continue
		}
	}

	return ans
}