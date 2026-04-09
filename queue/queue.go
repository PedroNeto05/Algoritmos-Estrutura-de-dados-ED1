package queue

type IQueue interface {
	Enqueue(val int)
	Dequeue() (int,error)
	Front() (int,error)
	IsEmpty() bool
	Size() int
}
