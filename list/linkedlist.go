package list

import "errors"

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head     *Node
	tail     *Node
	inserted int
}

func (l *LinkedList) Add(val int) {
	newNode := &Node{
		val:  val,
		next: nil,
	}

	currNode := l.head

	for currNode.next != nil {
		currNode = currNode.next
	}

	currNode.next = newNode
	l.tail = newNode

	l.inserted++
}

func (l *LinkedList) AddOnIndex(val, index int) error {
	if index >= l.inserted || index < 0 {
		return errors.New("index fora do vetor")
	}

	return nil
}

func (l *LinkedList) Get(index int) (int, error) {
	if l.inserted == 0 {
		return -1, errors.New("o vetor esta vazio")
	}
	if index >= l.inserted || index < 0 {
		return -1, errors.New("index fora do vetor")
	}

	count := 0
	currNode := l.head

	for currNode.next != nil {
		if count == index {
			break
		}
		currNode = currNode.next
		count++
	}
	return currNode.val, nil
}

func (l *LinkedList) Set(val, index int) error {
	if index >= l.inserted || index < 0 {
		return errors.New("index fora do vetor")
	}

	count := 0
	currNode := l.head

	for currNode.next != nil {
		if count == index {
			break
		}
		currNode = currNode.next
		count++
	}
	currNode.val = val
	return nil
}

// func (l *LinkedList) RemoveOnIndex(index int) error {}
func (l *LinkedList) Size() int {
	return l.inserted
}
