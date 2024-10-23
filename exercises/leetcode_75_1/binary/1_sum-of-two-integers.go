/*
	Problem: 371. Sum of Two Integers
	Reference: https://leetcode.com/problems/sum-of-two-integers/
	Complexity: O(log(n))
*/

func getSum(a int, b int) int {
    sum := a ^ b
    carry := (a & b) << 1
    if carry != 0 {
        return getSum(sum, carry)
    }
    return sum
}