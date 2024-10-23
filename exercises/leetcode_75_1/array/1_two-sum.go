/*
	Problem: 1. Two Sum
	Reference: https://leetcode.com/problems/two-sum/
	Complexity: O(n)
*/

func twoSum(nums []int, target int) []int {
    mapper := make(map[int]int)
    for i, n := range nums {
        if _, ok := mapper[n]; ok {
            return []int{mapper[n],i}
        }
        diff := target - n
        mapper[diff] = i
    }
    return []int{}
}
