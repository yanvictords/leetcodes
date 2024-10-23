/*
	Problem: 53. Maximum Subarray
	Reference: https://leetcode.com/problems/maximum-subarray/
	Algorithm: Kadane's algorithm
	Complexity: O(n)
*/

func maxSubArray(nums []int) int {
    sum, acc := -1000000, 0
    for _, n := range nums {
        if acc <= 0 {
            acc = 0
        }
        if acc + n > sum {
            sum = acc + n
        }
        acc += n
    }
    return sum
}
