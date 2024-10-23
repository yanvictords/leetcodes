/*
	Problem: 217. Contains Duplicate
	Reference: https://leetcode.com/problems/contains-duplicate/
	Complexity: O(n)
*/

func containsDuplicate(nums []int) bool {
    mapper := make(map[int]struct{})
    for _, n := range nums {
        if _, ok := mapper[n]; ok {
            return true
        }
        mapper[n] = struct{}{}
    }
    return false
}
