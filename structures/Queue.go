package structures

//queue is built upon a LinkedList
type queue struct {
	list *LinkedList
}

//NewQueue returns a reference to a new queue
func NewQueue() *queue {
	l := new(LinkedList)
	q := queue{list: l}
	return &q
}

func (q *queue) Enqueue(n *Node){
	q.list.Append(n)
}

func (q *queue) Dequeue() *Node{
	n := q.list.Root()
	q.list.RemoveAtStart()
	return n
}

func (q *queue) Peek() *Node{
	return q.list.Root()
}

//Empty returns true if the queue contains no elements
func (q *queue) Empty() bool {
	return q.list.Empty()
}

//Len returns the length of the queue
func (q *queue) Len() int {
	return q.list.Len()
}