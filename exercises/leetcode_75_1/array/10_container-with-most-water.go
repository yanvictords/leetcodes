/*
	Problem: 11. Container With Most Water
	Reference: https://leetcode.com/problems/container-with-most-water/
	Complexity: O(n)
*/

func calcArea(l, r, dist int) int {
    return min(l,r) * dist
}

func maxArea(height []int) int {
    total := 0
    l, r := 0, len(height) - 1
    for l < r {
        area := calcArea(height[l], height[r], r-l)
        total = max(total, area)
        if height[l] < height[r] {
            l++
            continue
        }
        r--
    }
    return total
}
