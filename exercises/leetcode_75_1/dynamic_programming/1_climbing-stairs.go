/*
	Problem: 70. Climbing Stairs
	Reference: https://leetcode.com/problems/climbing-stairs
	Complexity: O(N)
*/

func climb(n int, mapper map[int]int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if _, ok := mapper[n]; ok {
		return mapper[n]
	}

	mapper[n] = climb(n-1, mapper) + climb(n-2, mapper)
	return mapper[n]
}

func climbStairs(n int) int {
	mapper := make(map[int]int, n)
	return climb(n, mapper)
}

