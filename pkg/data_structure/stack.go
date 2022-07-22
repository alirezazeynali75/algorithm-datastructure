package datastructure

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(value string) {
	*s = append(*s, value)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		top := len(*s) - 1
		topElement := (*s)[top]
		*s = (*s)[:top]
		return topElement, true
	}
}
