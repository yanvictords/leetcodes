/*
	Problem: 56. Merge Intervals
	Reference: https://leetcode.com/problems/merge-intervals/
    Complexity: O(n log n)
*/

func merge(intervals [][]int) [][]int {
	ans := [][]int{}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	low, high := intervals[0][0], intervals[0][1]
	for idx, next := range intervals {
		if idx == 0 {
			continue
		}
		if high >= next[0] {
			low, high = min(low, next[0]), max(high, next[1])
			continue
		}
		ans = append(ans, []int{low, high})
		low, high = next[0], next[1]
	}
	return append(ans, []int{low, high})
}