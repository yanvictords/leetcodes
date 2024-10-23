/*
	Problem: 1143. Longest Common Subsequence
	Reference: https://leetcode.com/problems/longest-common-subsequence/
	Complexity: O(N * M)
*/

type Pair struct {
	I, J int
}

func dfs(text1 string, text2 string, pair Pair, answers map[Pair]int) int {
	if pair.I >= len(text1) || pair.J >= len(text2) {
		return 0
	}
	if ans, ok := answers[pair]; ok {
		return ans
	}

	answers[pair] = 0

	if text1[pair.I] == text2[pair.J] {
		answers[pair] = 1 + dfs(text1, text2, Pair{pair.I + 1, pair.J + 1}, answers)
		return answers[pair]
	}

	answers[pair] = max(
		dfs(text1, text2, Pair{pair.I + 1, pair.J}, answers),
		dfs(text1, text2, Pair{pair.I, pair.J + 1}, answers),
	)

	return answers[pair]
}

func longestCommonSubsequence(text1 string, text2 string) int {
	answers := make(map[Pair]int, len(text1)*len(text2))
	return dfs(text1, text2, Pair{0, 0}, answers)
}