/*
	Problem: 191. Number of 1 Bits
	Reference: https://leetcode.com/problems/number-of-1-bits
	Complexity: O(1)
*/

func hammingWeight(n int) int {
    total := 0
    for i := 0; i < 32; i++ {
        one := n & 1
        if one == 1 {
            total++
        }
        n = n >> 1
    }
    return total
}