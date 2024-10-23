/*
	Problem: 42. Trapping Rain Water
	Reference: https://leetcode.com/problems/trapping-rain-water
	Results: 10ms, Beats 65.81%
	Complexity: O(n)
*/

func calculateWater(height []int, l, r, midBlocks int) int {
	if height[l] < height[r] {
		ans := (height[l] * (r - 1 - l)) - midBlocks
		return ans
	}
	ans := (height[r] * (r - 1 - l)) - midBlocks
	return ans
}

func performToLeftEqual(height []int) int {
	r, water, midBlocks := len(height)-1, 0, 0
	for l := r - 1; l >= 0; l-- {
		if height[l] >= height[r] {
			water += calculateWater(height, l, r, midBlocks)
			midBlocks = 0
			r = l
			continue
		}
		midBlocks += height[l]
	}
	return water
}

func performToRight(height []int) int {
	l, water, midBlocks := 0, 0, 0
	for r := l + 1; r < len(height); r++ {
		if height[l] < height[r] {
			water += calculateWater(height, l, r, midBlocks)
			midBlocks = 0
			l = r
			continue
		}
		midBlocks += height[r]
	}
	return water
}

func trap(height []int) int {
	water := performToLeftEqual(height)
	water += performToRight(height)
	return water
}