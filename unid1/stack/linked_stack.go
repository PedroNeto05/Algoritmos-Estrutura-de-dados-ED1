package stack

import "errors"

type Node struct {
	val  int
	next *Node
}

type LinkedStack struct {
	top  *Node
	size int
}

func (s *LinkedStack) Push(value int) {
	node := &Node{
		val:  value,
		next: nil,
	}
	if s.size == 0 {
		s.top = node
		s.size++
		return
	}

	node.next = s.top
	s.top = node
	s.size++
}

func (s *LinkedStack) Pop() (int, error) {
	if s.size == 0 {
		return -1, errors.New("lista vazia")
	}

	val := s.top.val
	s.top = s.top.next
	s.size--
	return val, nil
}

func (s *LinkedStack) Peek() (int, error) {
	if s.size == 0 {
		return -1, errors.New("lista vazia")
	}

	return s.top.val, nil
}

func (s *LinkedStack) IsEmpty() bool {
	return s.size == 0
}

func (s *LinkedStack) Size() int {
	return s.size
}
