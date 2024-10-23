/*
	Problem: 5. Longest Palindromic Substring
	Reference: https://leetcode.com/problems/longest-palindromic-substring/
    Complexity: O(N^2)
*/

func longestPalindrome(s string) string {
	maxi := 0
	word := ""
	for i := range s {
		// impar
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > maxi {
				word = s[l : r+1]
				maxi = r - l + 1
			}
			l--
			r++
		}

		// par
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > maxi {
				word = s[l : r+1]
				maxi = r - l + 1
			}
			l--
			r++
		}

	}

	return word
}