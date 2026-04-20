package queue

import (
	"errors"
)

type ArrayQueue struct {
	arr  []int
	size int
	head int
	tail int
}

func (q *ArrayQueue) Init(size int) error {
	if size <= 0 {
		return errors.New("a fila não pode ter tamanho menor ou igual a zero")
	}

	q.arr = make([]int, size)

	return nil
}

func (q *ArrayQueue) duplicateQueue() {
	currSize := len(q.arr)
	newArr := make([]int, currSize*2)

	temp := q.head
	for i := 0; i < len(q.arr); i++ {
		newArr[i] = q.arr[temp]
		temp = (temp + 1) % len(q.arr)
	}
	q.head = 0
	q.tail = q.size

	q.arr = newArr
}

func (q *ArrayQueue) Enqueue(val int) {
	if q.size >= len(q.arr) {
		q.duplicateQueue()
	}

	q.arr[q.tail] = val
	q.tail = (q.tail + 1) % len(q.arr)
	q.size++
}

func (q *ArrayQueue) Dequeue() (int, error) {
	if q.size <= 0 {
		return -1, errors.New("lista vazia")
	}
	val := q.arr[q.head]

	q.head = (q.head + 1) % len(q.arr)

	q.size--
	return val, nil
}

func (q *ArrayQueue) Front() (int, error) {
	if q.size <= 0 {
		return -1, errors.New("lista vazia")
	}
	return q.arr[q.head], nil
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *ArrayQueue) Size() int {
	return q.size
}

func (q *ArrayQueue) SizeAlt() int {
	return (q.tail - q.head + len(q.arr)) % len(q.arr)
}
