/*
	Problem: 125. Valid Palindrome
	Reference: https://leetcode.com/problems/valid-palindrome/
    Complexity: O(N)
*/

func isPalindrome(s string) bool {
	word := ""
	for _, c := range s {
		if isAlphanumeric(c) {
			word += string(lower(c))
		}
	}
	for i, _ := range word {
		if word[i] != word[len(word)-1-i] {
			return false
		}
	}

	return true
}

func isAlphanumeric(c rune) bool {
	return c >= 'a' && c <= 'z' ||
		c >= 'A' && c <= 'Z' ||
		c >= '0' && c <= '9'
}

func lower(c rune) rune {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}