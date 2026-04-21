package deque

import "errors"

type Node struct {
	val  int
	next *Node
	prev *Node
}

type LinkedDeque struct {
	front *Node
	rear  *Node
	size  int
}

func (d *LinkedDeque) EnqueueFront(value int) {
	node := &Node{
		val:  value,
		next: nil,
		prev: nil,
	}

	if d.size == 0 {
		d.front = node
		d.rear = node
		d.size++
		return
	}

	d.front.prev = node
	node.next = d.front
	d.front = node
	d.size++
}

func (d *LinkedDeque) EnqueueRear(value int) {
	node := &Node{
		val:  value,
		next: nil,
		prev: nil,
	}

	if d.size == 0 {
		d.front = node
		d.rear = node
		d.size++
		return
	}

	d.rear.next = node
	node.prev = d.rear
	d.rear = node
	d.size++
}

func (d *LinkedDeque) DequeueFront() (int, error) {
	if d.size == 0 {
		return -1, errors.New("lista vazia")
	}

	val := d.front.val

	d.front = d.front.next

	if d.size == 1 {
		d.front = nil
		d.rear = nil
	}
	d.size--

	return val, nil
}

func (d *LinkedDeque) DequeueRear() (int, error) {
	if d.size == 0 {
		return -1, errors.New("lista vazia")
	}

	val := d.rear.val

	d.rear = d.rear.prev

	if d.size == 1 {
		d.front = nil
		d.rear = nil
	}
	d.size--

	return val, nil
}

func (d *LinkedDeque) Front() (int, error) {
	if d.size == 0 {
		return -1, errors.New("lista vazia")
	}
	return d.front.val, nil
}

func (d *LinkedDeque) Rear() (int, error) {
	if d.size == 0 {
		return -1, errors.New("lista vazia")
	}
	return d.rear.val, nil
}

func (d *LinkedDeque) IsEmpty() bool {
	return d.size == 0
}

func (d *LinkedDeque) Size() int {
	return d.size
}
