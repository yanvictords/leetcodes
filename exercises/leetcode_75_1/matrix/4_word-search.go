/*
	Problem: 79. Word Search
	Reference: https://leetcode.com/problems/word-search/
    Complexity: O(4^N)
*/

func exists(board [][]byte, word string, i, j int) bool {
	if len(word) == 0 {
		return true
	}
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) {
		return false
	}
	if word[0] == board[i][j] {
		temp := board[i][j]
		board[i][j] = '#'
		match := exists(board, word[1:], i+1, j) ||
			exists(board, word[1:], i, j+1) ||
			exists(board, word[1:], i-1, j) ||
			exists(board, word[1:], i, j-1)
		board[i][j] = temp
		return match
	}
	return false
}

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			if exists(board, word, i, j) {
				return true
			}
		}
	}
	return false
}