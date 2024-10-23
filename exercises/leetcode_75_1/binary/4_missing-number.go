/*
	Problem: 268. Missing Number
	Reference: https://leetcode.com/problems/missing-number
    Complexity: Time O(n), Memory O(1)
*/

func missingNumber(nums []int) int {
    max := len(nums)
    expected := (max * (max+1)) / 2
    sum := 0
    for _, n := range nums {
        sum += n
    }
 
    return expected - sum
}