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

func (q *LinkedQueue) Enqueue(val int) {
	node := &Node{
		val:  val,
		next: nil,
	}

	if q.size == 0 {
		q.front = node
		q.back = node
		q.size++
		return
	}

	q.back.next = node
	q.back = node
	q.size++
}

func (q *LinkedQueue) Dequeue() (int, error) {
	if q.size < 1 {
		return -1, errors.New("lista vazia")
	}
	val := q.front.val
	q.front = q.front.next
	q.size--
	if q.size == 0 {
		q.back = nil
	}
	return val, nil
}

func (q *LinkedQueue) Front() (int, error) {
	if q.size < 1 {
		return -1, errors.New("lista vazia")
	}
	return q.front.val, nil
}

func (q *LinkedQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *LinkedQueue) Size() int {
	return q.size
}
