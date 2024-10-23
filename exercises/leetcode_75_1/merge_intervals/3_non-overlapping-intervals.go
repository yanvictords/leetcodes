/*
	Problem: 435. Non-overlapping Intervals
	Reference: https://leetcode.com/problems/non-overlapping-intervals/
    Complexity: O(n log n)
*/

func nonIntervals(intervals [][]int) int {
	non := 1

	high := intervals[0][1]
	for idx, next := range intervals {
		if idx == 0 {
			continue
		}
		if high <= next[0] {
			non++
			high = next[1]
		}
	}

	return non
}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	return len(intervals) - nonIntervals(intervals)
}