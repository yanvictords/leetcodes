/*
	Problem: 55. Jump Game
	Reference: https://leetcode.com/problems/jump-game
	Complexity: O(N)
*/

func dfs(nums []int, idx int, answers map[int]int) int {
	if idx >= len(nums) {
		return 0
	}

	if ans, ok := answers[idx]; ok {
		return ans
	}

	answers[idx] = max(
		dfs(nums, idx+1, answers),
		nums[idx]+dfs(nums, idx+2, answers),
	)
	return answers[idx]
}

func rob(nums []int) int {
	answers := make(map[int]int)
	return dfs(nums, 0, answers)
}