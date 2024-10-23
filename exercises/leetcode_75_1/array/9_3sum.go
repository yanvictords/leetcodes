/*
	Problem: 15. 3Sum
	Reference: https://leetcode.com/problems/3sum/
	Complexity: O(Nˆ2)
*/

import (
	"sort"
)

type Triple struct {
	I, J, K int
}

func toTriple(nums []int) Triple {
	sort.Ints(nums)
	return Triple{nums[0], nums[1], nums[2]}
}

func twoSum(nums []int, target int) [][]int {
	ans := [][]int{}
	mapper := make(map[int]struct{})
	for _, num := range nums {
		complement := (num + target) * (-1)
		if _, ok := mapper[complement]; ok {
			ans = append(ans, []int{num, complement})
		}
		mapper[num] = struct{}{}
	}
	return ans
}

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	answers := make(map[Triple]bool)
	for i := 0; i < len(nums); i++ {
		complements := twoSum(nums[i+1:], nums[i])
		for _, c := range complements {
			sec, thir := c[0], c[1]
			triple := toTriple([]int{nums[i], sec, thir})
			if !answers[triple] {
				ans = append(ans, []int{nums[i], sec, thir})
				answers[triple] = true
			}
		}
	}
	return ans
}