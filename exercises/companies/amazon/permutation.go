/*
	Problem: 46. Permutations
	Reference: https://leetcode.com/problems/permutations/
	Results: 0 ms, faster than 100.00%
	Complexity: O(1)
*/

func permute(nums []int) [][]int {
	if len(nums) < 2 {
		return [][]int{nums}
	}
	permutations := [][]int{}
	for i, num := range nums {
		complement := make([]int, len(nums)-1)
		copy(complement[:i], nums[:i])
		copy(complement[i:], nums[i+1:])
		subpermutations := permute(complement)
		for _, subp := range subpermutations {
			permutations = append(permutations, append([]int{num}, subp...))
		}
	}

	return permutations
}