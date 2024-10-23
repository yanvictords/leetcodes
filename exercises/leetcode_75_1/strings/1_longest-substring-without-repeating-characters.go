/*
	Problem: 3. Longest Substring Without Repeating Characters
	Reference: https://leetcode.com/problems/longest-substring-without-repeating-characters/
    Complexity: O(N)
*/

func lengthOfLongestSubstring(s string) int {
	length := 0
	mapper := make(map[byte]int)
	l := 0
	for r, _ := range s {
		if repeated, ok := mapper[s[r]]; ok {
			for l <= repeated {
				delete(mapper, byte(s[l]))
				l++
			}
		}
		mapper[s[r]] = r
		length = max(length, r-l+1)
	}

	return length
}