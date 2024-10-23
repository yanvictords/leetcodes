/*
	Problem: 128. Longest Consecutive Sequence
	Reference: https://leetcode.com/problems/longest-consecutive-sequence
    Complexity: O(n)
*/

func longestConsecutive(nums []int) int {
	rootmap := make(map[int]bool)
	subarr := make(map[int]*int)
	for _, num := range nums {
		if rootmap[num-1] {
			subarr[num-1] = &num
		}
		if rootmap[num+1] {
			next := num + 1
			subarr[num] = &next
		}
		rootmap[num] = true
	}
	max := 0
	for _, num := range subarr {
		delete(rootmap, *num)
	}
	for key, _ := range rootmap {
		size := 1
		next := subarr[key]
		for next != nil {
			size++
			next = subarr[*next]
		}
		if max < size {
			max = size
		}
	}
	return max
}
