/*
	Problem: 57. Insert Interval
	Reference: https://leetcode.com/problems/insert-interval/
    Complexity: O(n log n)
*/

func merge(intervals [][]int) [][]int {
	ans := [][]int{}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	low, high := 0, 0
	for idx, next := range intervals {
		if idx == 0 {
			low, high = intervals[0][0], intervals[0][1]
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

func insert(intervals [][]int, newInterval []int) [][]int {
	return merge(append(intervals, newInterval))
}