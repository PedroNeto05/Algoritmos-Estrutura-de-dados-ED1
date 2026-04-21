package deque

import "errors"

type ArrayDeque struct {
	arr  []int
	head int
	rear int
	size int
}

func (d *ArrayDeque) Init(size int) error {
	if size <= 0 {
		return errors.New("a fila não pode ter tamanho menor ou igual a zero")
	}

	d.arr = make([]int, size)
	d.head = 0
	d.rear = 0
	d.size = 0

	return nil
}

func (d *ArrayDeque) duplicateDeque() {
	currSize := len(d.arr)
	newArr := make([]int, currSize*2)

	temp := d.head
	for i := 0; i < d.size; i++ {
		newArr[i] = d.arr[temp]
		temp = (temp + 1) % len(d.arr)
	}

	d.head = 0
	d.rear = d.size

	d.arr = newArr
}

func (d *ArrayDeque) EnqueueFront(value int) {
	if d.size >= len(d.arr) {
		d.duplicateDeque()
	}
	d.head = (d.head - 1 + len(d.arr)) % len(d.arr)
	d.arr[d.head] = value
	d.size++
}

func (d *ArrayDeque) EnqueueRear(value int) {
	if d.size >= len(d.arr) {
		d.duplicateDeque()
	}

	d.arr[d.rear] = value
	d.rear = (d.rear + 1) % len(d.arr)
	d.size++
}

func (d *ArrayDeque) DequeueFront() (int, error) {
	if d.size <= 0 {
		return -1, errors.New("lista vazia")
	}
	val := d.arr[d.head]
	d.head = (d.head + 1) % len(d.arr)
	d.size--
	return val, nil
}

func (d *ArrayDeque) DequeueRear() (int, error) {
	if d.size <= 0 {
		return -1, errors.New("lista vazia")
	}
	index := (d.rear - 1 + len(d.arr)) % len(d.arr)
	val := d.arr[index]
	d.rear = index
	d.size--
	return val, nil
}

func (d *ArrayDeque) Front() (int, error) {
	if d.size <= 0 {
		return -1, errors.New("lista vazia")
	}
	return d.arr[d.head], nil
}

func (d *ArrayDeque) Rear() (int, error) {
	if d.size <= 0 {
		return -1, errors.New("lista vazia")
	}

	index := (d.rear - 1 + len(d.arr)) % len(d.arr)
	return d.arr[index], nil
}

func (d *ArrayDeque) IsEmpty() bool {
	return d.size == 0
}

func (d *ArrayDeque) Size() int {
	return d.size
}
