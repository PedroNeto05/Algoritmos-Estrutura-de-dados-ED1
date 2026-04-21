package deque

type IDeque interface {
	EnqueueFront(value int)
	EnqueueRear(value int)
	DequeueFront() (int, error)
	DequeueRear() (int, error)
	Front() (int, error)
	Rear() (int, error)
	IsEmpty() bool
	Size() int
}
