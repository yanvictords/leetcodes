/*
	Problem: 211. Design Add and Search Words Data Structure
	Reference: https://leetcode.com/problems/design-add-and-search-words-data-structure/
    Complexity: O(M * 26^N)
*/

type Letter struct {
	catalog map[byte]*Letter
	word    string
}

func NewLetter() *Letter {
	return &Letter{
		catalog: make(map[byte]*Letter, 26),
		word:    "",
	}
}

func (l *Letter) Insert(letter byte, word string) *Letter {
	if l.catalog[letter] == nil {
		l.catalog[letter] = NewLetter()
	}
	if !l.catalog[letter].HasWord() {
		l.catalog[letter].word = word
	}
	return l
}

func (l *Letter) Get(letter byte) *Letter {
	return l.catalog[letter]
}

func (l *Letter) GetAll() map[byte]*Letter {
	return l.catalog
}

func (l *Letter) HasWord() bool {
	return l.word != ""
}

type WordDictionary struct {
	trie *Letter
}

func Constructor() WordDictionary {
	return WordDictionary{
		trie: NewLetter(),
	}
}

func (this *WordDictionary) AddWord(word string) {
	letter := this.trie
	for i := 0; i < len(word)-1; i++ {
		w := word[i]
		letter = letter.Insert(w, "").Get(w)
	}
	w := word[len(word)-1]
	letter.Insert(w, word)
}

func (this *WordDictionary) Search(word string) bool {
	return this.searchAll(this.trie, word)
}

func (this *WordDictionary) searchAll(letter *Letter, word string) bool {
	for idx, char := range word {
		if char != '.' {
			if letter = letter.Get(byte(char)); letter == nil {
				return false
			}
			continue
		}
		for _, dotLetter := range letter.GetAll() {
			if this.searchAll(dotLetter, word[idx+1:]) {
				return true
			}
		}
		return false
	}

	return letter.HasWord()
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */