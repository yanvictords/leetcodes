/*
	Problem: 208. Implement Trie (Prefix Tree)
	Reference: https://leetcode.com/problems/implement-trie-prefix-tree
    Complexity: O(N)
*/

type Letter struct {
	catalog map[byte]*Letter
	word    string
}

func (l *Letter) Insert(letter byte, word string) *Letter {
	if _, ok := l.catalog[letter]; !ok {
		l.catalog[letter] = &Letter{
			catalog: make(map[byte]*Letter, 26),
		}
	}
	if !l.catalog[letter].HasWord() {
		l.catalog[letter].word = word
	}
	return l
}

func (l *Letter) Get(letter byte) *Letter {
	return l.catalog[letter]
}

func (l *Letter) HasWord() bool {
	return l.word != ""
}

type Trie struct {
	Head *Letter
}

func Constructor() Trie {
	return Trie{
		Head: &Letter{
			catalog: make(map[byte]*Letter, 26),
			word:    "",
		},
	}
}

func (this *Trie) Insert(word string) {
	letter := this.Head
	for i := 0; i < len(word)-1; i++ {
		w := word[i]
		letter = letter.Insert(w, "").Get(w)
	}
	w := word[len(word)-1]
	letter.Insert(w, word)
}

func (this *Trie) Search(word string) bool {
	letter := this.Head
	for _, w := range word {
		if letter = letter.Get(byte(w)); letter == nil {
			return false
		}
	}
	return letter.HasWord()
}

func (this *Trie) StartsWith(prefix string) bool {
	letter := this.Head
	for _, p := range prefix {
		if letter = letter.Get(byte(p)); letter == nil {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */