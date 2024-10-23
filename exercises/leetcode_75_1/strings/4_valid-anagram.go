/*
	Problem: 242. Valid Anagram
	Reference: https://leetcode.com/problems/valid-anagram/
    Complexity: O(N)
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapper := make(map[byte]int)
	for i, _ := range s {
		mapper[t[i]]++
		if mapper[t[i]] == 0 {
			delete(mapper, t[i])
		}
		mapper[s[i]]--
		if mapper[s[i]] == 0 {
			delete(mapper, s[i])
		}
	}
	return len(mapper) == 0
}