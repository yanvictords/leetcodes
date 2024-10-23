/*
	Problem: 49. Group Anagrams
	Reference: https://leetcode.com/problems/group-anagrams/description/
    Complexity: O(N*M*Log(M))
*/

func groupAnagrams(strs []string) [][]string {
	ans := [][]string{}
	mapper := make(map[string][]string)
	for _, s := range strs {
		sorted := sort(s)
		mapper[sorted] = append(mapper[sorted], s)
	}

	for _, s := range mapper {
		ans = append(ans, s)
	}
	return ans
}

func sort(str string) string {
	if len(str) < 2 {
		return str
	}
	sz := len(str)
	str1 := sort(str[sz/2:])
	str2 := sort(str[:sz/2])
	return merge(str1, str2)
}

func merge(str1 string, str2 string) string {
	ans := ""
	i, j := 0, 0
	for i < len(str1) && j < len(str2) {
		if str1[i] < str2[j] {
			ans += string(str1[i])
			i++
			continue
		}
		ans += string(str2[j])
		j++
	}

	for ; i < len(str1); i++ {
		ans += string(str1[i])
	}
	for ; j < len(str2); j++ {
		ans += string(str2[j])
	}
	return ans
}