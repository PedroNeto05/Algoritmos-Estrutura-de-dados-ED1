package queue

import "errors"

type Node struct {
	val  int
	next *Node
}

type LinkedQueue struct {
	front *Node
	back  *Node
	size  int
}

func (l *LinkedQueue) Enqueue(val int) {
	node := &Node{
		val:  val,
		next: nil,
	}

	if l.size == 0 {
		l.front = node
		l.back = node
		l.size++
		return
	}

	l.back.next = node
	l.back = node
	l.size++
}

func (l *LinkedQueue) Dequeue() (int, error) {
	if l.size < 1 {
		return -1, errors.New("lista vazia")
	}
	val := l.front.val
	l.front = l.front.next
	l.size--
	if l.size == 0 {
		l.back = nil
	}
	return val, nil
}

func (l *LinkedQueue) Front() (int, error) {
	if l.size < 1 {
		return -1, errors.New("lista vazia")
	}
	return l.front.val, nil
}

func (l *LinkedQueue) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedQueue) Size() int {
	return l.size
}

