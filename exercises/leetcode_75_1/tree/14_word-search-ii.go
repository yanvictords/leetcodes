/*
	Problem: 212. Word Search II
	Reference: https://leetcode.com/problems/word-search-ii/
    Complexity: O(M*N*4^K)
*/

type Trie struct {
	head *Letter
}

func NewTrie() *Trie {
	return &Trie{
		head: NewLetter(),
	}
}

func (t *Trie) Insert(word string) {
	letter := t.head
	for i := 0; i < len(word)-1; i++ {
		w := word[i]
		letter = letter.Insert(w, "").Get(w)
	}
	w := word[len(word)-1]
	letter.Insert(w, word)
}

func (t *Trie) Head() *Letter {
	return t.head
}

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
		l.catalog[letter] = NewLetter()
	}
	if !l.catalog[letter].HasWord() {
		l.catalog[letter].word = word
	}
	return l
}

func (l *Letter) Remove(letter byte) string {
	word := l.catalog[letter].word
	l.catalog[letter].word = ""
	return word
}

func (l *Letter) Get(next byte) *Letter {
	return l.catalog[next]
}

func (l *Letter) HasWord() bool {
	return l.word != ""
}

func clearTrie(letter *Letter, board [][]byte, i, j int, ans *[]string) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) {
		return
	}
	cell := board[i][j]
	next := letter.Get(cell)
	if next == nil {
		return
	}
	if next.HasWord() {
		*ans = append(*ans, letter.Remove(cell))
	}
	board[i][j] = '#'
	clearTrie(next, board, i-1, j, ans)
	clearTrie(next, board, i+1, j, ans)
	clearTrie(next, board, i, j-1, ans)
	clearTrie(next, board, i, j+1, ans)
	board[i][j] = cell
}

func findWords(board [][]byte, words []string) []string {
	ans := []string{}
	trie := NewTrie()
	for _, word := range words {
		trie.Insert(word)
	}
	for i := range board {
		for j := range board[i] {
			clearTrie(trie.Head(), board, i, j, &ans)
		}
	}
	return ans
}