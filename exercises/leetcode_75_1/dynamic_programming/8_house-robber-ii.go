/*
	Problem: 213. House Robber II
	Reference: https://leetcode.com/problems/house-robber-ii/
	Complexity: O(2^N)
*/

type Pair struct {
	Idx           int
	IncludesFirst bool
}

func dfs(nums []int, idx int, answers map[Pair]int, includesFirst bool) int {
	if idx >= len(nums) {
		return 0
	}
	if idx == len(nums)-1 && includesFirst {
		return 0
	}

	pair := Pair{idx, includesFirst}
	if ans, ok := answers[pair]; ok {
		return ans
	}

	answers[pair] = max(
		dfs(nums, idx+1, answers, includesFirst),
		nums[idx]+dfs(nums, idx+2, answers, idx == 0 || includesFirst),
	)
	return answers[pair]
}

func rob(nums []int) int {
	answers := make(map[Pair]int)
	return dfs(nums, 0, answers, false)
}