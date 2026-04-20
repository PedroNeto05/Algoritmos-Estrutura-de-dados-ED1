package list

import "errors"

type ArrayList struct {
	arr      []int
	inserted int
}

func (l *ArrayList) Init(size int) error {
	if size <= 0 {
		return errors.New("a lista não pode ter tamanho menor ou igual a zero")
	}

	l.arr = make([]int, size)

	return nil
}

func (l *ArrayList) duplicateList() {
	currSize := len(l.arr)
	newArr := make([]int, currSize*2)

	for i := 0; i < l.inserted; i++ {
		newArr[i] = l.arr[i]
	}

	l.arr = newArr
}

func (l *ArrayList) Add(val int) {
	if l.inserted >= len(l.arr) {
		l.duplicateList()
	}

	l.arr[l.inserted] = val
	l.inserted++
}

func (l *ArrayList) AddOnIndex(val, index int) error {
	if index > l.inserted || index < 0 {
		return errors.New("index fora do vetor")
	}
	if l.inserted >= len(l.arr) {
		l.duplicateList()
	}

	for i := l.inserted; i > index; i-- {
		l.arr[i] = l.arr[i-1]
	}

	l.arr[index] = val
	l.inserted++

	return nil
}

func (l *ArrayList) Get(index int) (int, error) {
	if l.inserted == 0 {
		return -1, errors.New("o vetor esta vazio")
	}
	if index >= l.inserted || index < 0 {
		return -1, errors.New("index fora do vetor")
	}
	return l.arr[index], nil
}

func (l *ArrayList) Set(val, index int) error {
	if l.inserted == 0 {
		return errors.New("o vetor esta vazio")
	}
	if index >= l.inserted || index < 0 {
		return errors.New("index fora do vetor")
	}

	l.arr[index] = val
	return nil
}

func (l *ArrayList) RemoveOnIndex(index int) error {
	if l.inserted == 0 {
		return errors.New("o vetor esta vazio")
	}
	if index >= l.inserted || index < 0 {
		return errors.New("index fora do vetor")
	}

	for i := index; i < l.inserted-1; i++ {
		l.arr[i] = l.arr[i+1]
	}
	l.inserted--
	l.arr[l.inserted] = 0
	return nil
}

func (l *ArrayList) Size() int {
	return l.inserted
}

func (l *ArrayList) Reverse() {
	limit := l.inserted / 2
	for i := range limit {
		l.arr[i], l.arr[l.inserted-1-i] = l.arr[l.inserted-1-i], l.arr[i]
	}
}
