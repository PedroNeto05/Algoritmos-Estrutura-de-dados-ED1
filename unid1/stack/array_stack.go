package stack

import "errors"

type ArrayStack struct {
	arr  []int
	size int
}

func (s *ArrayStack) Init(size int) error {
	if size <= 0 {
		return errors.New("a pilha não pode ter tamanho menor ou igual a zero")
	}

	s.arr = make([]int, size)

	return nil
}

func (s *ArrayStack) duplicateStack() {
	currSize := len(s.arr)
	newArr := make([]int, currSize*2)

	for i := 0; i < s.size; i++ {
		newArr[i] = s.arr[i]
	}

	s.arr = newArr
}

func (s *ArrayStack) Push(value int) {
	if s.size >= len(s.arr) {
		s.duplicateStack()
	}
	s.arr[s.size] = value
	s.size++
}

func (s *ArrayStack) Pop() (int, error) {
	if s.size <= 0 {
		return -1, errors.New("lista vazia")
	}
	s.size--
	return s.arr[s.size], nil
}

func (s *ArrayStack) Peek() (int, error) {
	if s.size <= 0 {
		return -1, errors.New("lista vazia")
	}
	return s.arr[s.size-1], nil
}

func (s *ArrayStack) IsEmpty() bool {
	return s.size == 0
}

func (s *ArrayStack) Size() int {
	return s.size
}
