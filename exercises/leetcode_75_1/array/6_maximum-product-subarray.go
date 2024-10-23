/*
	Problem: 152. Maximum Product Subarray
	Reference: https://leetcode.com/problems/maximum-product-subarray/
	Complexity: O(n)
*/

func maxProduct(nums []int) int {
    n := len(nums)
    curMax, curMin := 1, 1

    res := math.MinInt
    for i := 0; i < n; i++ {
        prodMax := curMax * nums[i]
        prodMin := curMin * nums[i]
        curMax = max(nums[i], prodMin, prodMax)
        curMin = min(nums[i], prodMin, prodMax)
        res = max(res, curMax)
    }
    return res
}
