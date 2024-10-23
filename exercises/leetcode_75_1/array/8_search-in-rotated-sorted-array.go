/*
	Problem: 33. Search in Rotated Sorted Array
	Reference: https://leetcode.com/problems/search-in-rotated-sorted-array/
	Complexity: O(log(n))
*/

func findPivot(nums []int, l, h int) int {
    mid := (l + h) / 2
    if l > h || mid == len(nums) - 1 {
        return len(nums) - 1
    }
    if nums[mid] > nums[mid+1] {
        return mid
    }
    if nums[mid] > nums[len(nums) - 1] {
        return findPivot(nums, mid + 1, h)
    }
    return findPivot(nums, l, mid - 1)
}

func binarySearch(nums []int, target, l, h int) int {
    if l > h {
        return -1
    }
    mid := (l + h) / 2
    if nums[mid] == target {
        return mid
    }
    if target > nums[mid] {
        return binarySearch(nums, target, mid + 1, h)
    }
    return binarySearch(nums, target, l, mid - 1)
}

func search(nums []int, target int) int {
    pivot := findPivot(nums, 0, len(nums) - 1)
    if target >= nums[0] && target <= nums[pivot] {
        return binarySearch(nums, target, 0, pivot)
    }
    return binarySearch(nums, target, pivot + 1, len(nums) - 1)
}
