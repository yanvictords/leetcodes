/*
	Problem: 153. Find Minimum in Rotated Sorted
	Reference: https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
	Complexity: O(log(n))
*/

func binarySearch(nums []int, l, h int) int {
    mid := (l + h) / 2
    if l > h || mid == (len(nums) - 1) {
        return nums[0]
    }
    if nums[mid] > nums[mid+1] {
        return nums[mid+1]
    }
    if nums[mid] > nums[len(nums)-1] {
        return binarySearch(nums, mid + 1, h)
    }
    return binarySearch(nums, l, mid - 1)
}

func findMin(nums []int) int {
    return binarySearch(nums, 0, len(nums) - 1)
}
