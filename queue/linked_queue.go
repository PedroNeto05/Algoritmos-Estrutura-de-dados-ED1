package queue 

type Node struct {
	val int
	next *Node
}

type LinkedQueue struct {
	front *Node
	back *Node
	size int
}

func (l *LinkedQueue) Enqueue(val int) {}
func (l *LinkedQueue) Dequeue() (int,error) {
	if l.size < 1 {
		return -1,errors.New("lista vazia")
	}
}

func (l *LinkedQueue) Front() (int,error) {
	if l.size < 1 {
		return -1,errors.New("lista vazia")
	}
	return l.front.val,nil
}
func (l *LinkedQueue) IsEmpty() bool {
	return l.size == 0
}
func (l *LinkedQueue) Size() int {
	return l.size
}