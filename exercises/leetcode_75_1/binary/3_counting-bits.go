/*
	Problem: 338. Counting Bits
	Reference: https://leetcode.com/problems/counting-bits/
	Complexity: O(n)
*/

func count(n int) int {
    ans := 0
    for i := 0; i < 32; i++ {
        one := n & 1
        if one == 1 {
            ans++
        }
        n = n >> 1
    }
    return ans
}

func countBits(n int) []int {
    ans := []int{}
    for i := 0; i <= n; i++ {
        ans = append(ans, count(i))
    }
    return ans
}