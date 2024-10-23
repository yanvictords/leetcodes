/*
	Problem: 377. Combination Sum IV
	Reference: https://leetcode.com/problems/combination-sum-iv
	Complexity: O(N * Target)
*/

func dfs(nums []int, target int, answers map[int]int) int {
	if target == 0 {
		return 1
	}

	if ans, ok := answers[target]; ok {
		return ans
	}

	answers[target] = 0
	for _, num := range nums {
		if target >= num {
			answers[target] += dfs(nums, target-num, answers)
		}
	}

	return answers[target]
}

func combinationSum4(nums []int, target int) int {
	answers := make(map[int]int)
	return dfs(nums, target, answers)
}

