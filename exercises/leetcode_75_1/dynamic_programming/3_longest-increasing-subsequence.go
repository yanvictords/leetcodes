/*
	Problem: 300. Longest Increasing Subsequence
	Reference: https://leetcode.com/problems/longest-increasing-subsequence/
	Complexity: O(N^2)
*/

func dfs(nums []int, curr int, answers map[int]int) int {
	if curr == len(nums)-1 {
		answers[curr] = 1
	}
	if val, ok := answers[curr]; ok {
		return val
	}

	answers[curr] = 1
	for next := curr + 1; next < len(nums); next++ {
		if nums[curr] < nums[next] {
			answers[curr] = max(answers[curr], 1+dfs(nums, next, answers))
		}
	}

	return answers[curr]
}

func lengthOfLIS(nums []int) int {
	answers := make(map[int]int, len(nums))
	ans := 0
	for curr := range nums {
		answers[curr] = dfs(nums, curr, answers)
	}
	for _, val := range answers {
		if val > ans {
			ans = val
		}
	}
	return ans
}