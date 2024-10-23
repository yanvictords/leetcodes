/*
	Problem: 190. Reverse Bits
	Reference: https://leetcode.com/problems/reverse-bits/
	Results: 0ms, Beats 100.00%
	Complexity: O(1)
*/

func reverseBits(num uint32) uint32 {
    ans := uint32(0)
    for i := 0; i < 32; i++ {
        bit := (num >> i) & 1
        ans += bit << (31 - i)
    }
    return ans
}