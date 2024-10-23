/*
	Problem: 238. Product of Array Except Self
	Reference: https://leetcode.com/problems/product-of-array-except-self/
	Complexity: O(n)
*/

func productExceptSelf(nums []int) []int {
    sz := len(nums)
    ans := make([]int, sz)
    for i := range nums {
        ans[i] = 1
    }
    l,r := 1, 1
    for i := range nums {
        ans[i] *= l
        ans[sz-i-1] *= r
        l *= nums[i]
        r *= nums[sz-1-i]
    }
    return ans
}
