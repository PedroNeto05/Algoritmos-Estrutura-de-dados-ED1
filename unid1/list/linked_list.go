package list

import (
	"errors"
)

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head     *Node
	inserted int
}

func (l *LinkedList) Add(val int) {
	newNode := &Node{
		val:  val,
		next: nil,
	}

	if l.inserted == 0 {
		l.head = newNode
		l.inserted++
		return
	}

	currNode := l.head

	for currNode.next != nil {
		currNode = currNode.next
	}

	currNode.next = newNode

	l.inserted++
}

func (l *LinkedList) AddOnIndex(val, index int) error {
	if index > l.inserted || index < 0 {
		return errors.New("index fora do vetor")
	}
	if l.inserted == 0 && index > 0 {
		return errors.New("index fora do vetor")
	}

	prev := l.head
	newNode := &Node{
		val:  val,
		next: nil,
	}

	if index == 0 {
		newNode.next = l.head
		l.head = newNode
		l.inserted++
		return nil
	}

	for i := 0; i < index-1; i++ {
		prev = prev.next
	}

	newNode.next = prev.next
	prev.next = newNode
	l.inserted++

	return nil
}

func (l *LinkedList) Get(index int) (int, error) {
	if l.inserted == 0 {
		return -1, errors.New("o vetor esta vazio")
	}
	if index >= l.inserted || index < 0 {
		return -1, errors.New("index fora do vetor")
	}

	currNode := l.head

	for range index {
		currNode = currNode.next
	}

	return currNode.val, nil
}

func (l *LinkedList) Set(val, index int) error {
	if index >= l.inserted || index < 0 {
		return errors.New("index fora do vetor")
	}

	currNode := l.head

	for range index {
		currNode = currNode.next
	}

	currNode.val = val
	return nil
}

func (l *LinkedList) RemoveOnIndex(index int) error {
	if index >= l.inserted || index < 0 {
		return errors.New("index fora do vetor")
	}

	if index == 0 {
		l.head = l.head.next
		l.inserted--
		return nil
	}

	prev := l.head

	for i := 0; i < index-1; i++ {
		prev = prev.next
	}

	prev.next = prev.next.next
	l.inserted--

	return nil
}

func (l *LinkedList) Size() int {
	return l.inserted
}

func (l *LinkedList) Reverse() {
	var prev *Node = nil

	curr := l.head

	var next *Node = nil

	for curr != nil {
		next = curr.next

		curr.next = prev

		prev = curr
		curr = next
	}
	l.head = prev
}
