package algoritmos

import "algoritmosestruturadados/stack"

func IsValidParenteses(target string) bool {
	s := &stack.LinkedStack{}

	for _, v := range target {
		if v == '(' {
			s.Push(1)
			continue
		}
		if _, err := s.Pop(); err != nil {
			return false
		}
	}

	return s.IsEmpty()
}
