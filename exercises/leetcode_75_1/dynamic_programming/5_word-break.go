/*
	Problem: 139. Word Break
	Reference: https://leetcode.com/problems/word-break/
	Complexity: O(N * M)
*/

type Letter struct {
	catalog map[byte]*Letter
	word    string
}

func NewLetter() *Letter {
	return &Letter{
		catalog: make(map[byte]*Letter),
		word:    "",
	}
}

func (l *Letter) Insert(letter byte, word string) *Letter {
	if l.catalog[letter] == nil {
		l.catalog[letter] = &Letter{
			catalog: make(map[byte]*Letter, 26),
		}
	}
	if word != "" {
		l.catalog[letter].word = word
	}
	return l
}

func (l *Letter) Get(letter byte) *Letter {
	return l.catalog[letter]
}

func (l *Letter) Match(word string) bool {
	return l.word == word
}

type Trie struct {
	root *Letter
}

func NewTrie() *Trie {
	return &Trie{
		root: NewLetter(),
	}
}

func (t *Trie) Insert(word string) {
	letter := t.root
	for i := 0; i < len(word)-1; i++ {
		w := word[i]
		letter = letter.Insert(w, "").Get(w)
	}
	w := word[len(word)-1]
	letter.Insert(w, word)
}

func (t *Trie) Match(word string) bool {
	letter := t.root
	for _, w := range word {
		letter = letter.Get(byte(w))
		if letter == nil {
			return false
		}
	}

	return letter.Match(word)
}

func dfs(s string, wordDict *Trie, answers map[string]bool) bool {
	if s == "" {
		return true
	}

	if ans, ok := answers[s]; ok {
		return ans
	}

	answers[s] = false
	word := ""
	for idx, c := range s {
		word += string(c)
		if wordDict.Match(word) {
			answers[s] = answers[s] || dfs(s[idx+1:], wordDict, answers)
		}
	}

	return answers[s]
}

func wordBreak(s string, wordDict []string) bool {
	trie := NewTrie()
	answers := make(map[string]bool)

	for _, wd := range wordDict {
		trie.Insert(wd)
	}

	return dfs(s, trie, answers)
}