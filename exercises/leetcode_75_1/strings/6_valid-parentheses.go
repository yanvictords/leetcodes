/*
	Problem: 20. Valid Parentheses
	Reference: https://leetcode.com/problems/valid-parentheses/
    Complexity: O(N)
*/

type Stack struct {
	parenthesis []rune
}

func (s *Stack) Empty() bool {
	return len(s.parenthesis) == 0
}

func (s *Stack) Insert(top rune) {
	if s.parenthesis == nil {
		s.parenthesis = []rune{}
	}
	s.parenthesis = append(s.parenthesis, top)
}

func (s *Stack) Pop() rune {
	if s.Empty() {
		return ' '
	}
	top := s.parenthesis[len(s.parenthesis)-1]
	s.parenthesis = s.parenthesis[:len(s.parenthesis)-1]
	return top
}

func isValid(s string) bool {
	stack := Stack{}
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack.Insert(c)
			continue
		}
		top := stack.Pop()
		if top != complement(c) {
			return false
		}
	}

	return stack.Empty()
}

func complement(close rune) rune {
	if close == ')' {
		return '('
	}
	if close == '}' {
		return '{'
	}
	return '['
}