package list

import "errors"

type Nodes struct {
	val  int
	next *Nodes
	prev *Nodes
}
type DoublyLinkedList struct {
	head     *Nodes
	tail     *Nodes
	inserted int
}

func (l *DoublyLinkedList) Add(val int) {
	node := &Nodes{
		val:  val,
		next: nil,
		prev: nil,
	}

	if l.inserted == 0 {
		l.head = node
		l.tail = node
		l.inserted++
		return
	}

	node.prev = l.tail
	l.tail.next = node
	l.tail = node
	l.inserted++
}

func (l *DoublyLinkedList) AddOnIndex(val, index int) error {
	if index > l.inserted || index < 0 {
		return errors.New("index fora do vetor")
	}

	node := &Nodes{
		val:  val,
		next: nil,
		prev: nil,
	}
	if l.inserted == 0 {
		l.head = node
		l.tail = node
		l.inserted++
		return nil
	}
	if index == 0 {
		l.head.prev = node
		node.next = l.head
		l.head = node
		l.inserted++
		return nil
	}
	if l.inserted == index {
		node.prev = l.tail
		l.tail.next = node
		l.tail = node
		l.inserted++
		return nil
	}
	currNodes := l.head
	for range index {
		currNodes = currNodes.next
	}
	node.next = currNodes
	node.prev = currNodes.prev
	currNodes.prev.next = node
	currNodes.prev = node
	l.inserted++
	return nil
}

func (l *DoublyLinkedList) Get(index int) (int, error) {
	if index >= l.inserted || index < 0 {
		return -1, errors.New("index fora do vetor")
	}

	currNodes := l.head
	for range index {
		currNodes = currNodes.next
	}
	return currNodes.val, nil
}

func (l *DoublyLinkedList) Set(val, index int) error {
	if index >= l.inserted || index < 0 {
		return errors.New("index fora do vetor")
	}

	currNodes := l.head
	for range index {
		currNodes = currNodes.next
	}
	currNodes.val = val
	return nil
}

func (l *DoublyLinkedList) RemoveOnIndex(index int) error {
	if index >= l.inserted || index < 0 {
		return errors.New("index fora do vetor")
	}

	if index == 0 {
		l.head = l.head.next
		if l.head != nil {
			l.head.prev = nil
		} else {
			l.tail = nil
		}
		l.inserted--
		return nil
	}

	currNodes := l.head
	for range index {
		currNodes = currNodes.next
	}

	if currNodes.next == nil {
		l.tail = currNodes.prev
		l.tail.next = nil
		l.inserted--
		return nil
	}

	currNodes.prev.next = currNodes.next

	currNodes.next.prev = currNodes.prev

	l.inserted--
	return nil
}

func (l *DoublyLinkedList) Size() int {
	return l.inserted
}
