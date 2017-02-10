package structures

//stack is built upon a LinkedList
type stack struct {
	list *LinkedList
}

//Push adds a Node to the stack
func (s *stack) Push(n *Node) {
	s.list.Append(n)
}

//NewStack returns a reference to a new stack
func NewStack() *stack {
	l := new(LinkedList)
	s := stack{list: l}
	return &s
}

//Pop removes and returns the head Node from the stack
func (s *stack) Pop() *Node {
	n := s.Peek()
	s.list.RemoveAtEnd()
	return n
}

//Peek returns the head Node from the stack
func (s *stack) Peek() *Node {
	return s.list.Peek()
}

//Empty returns true if the stack contains no elements
func (s *stack) Empty() bool {
	return s.list.Empty()
}

//Depth returns the depth(length) of the stack
func (s *stack) Depth() int {
	return s.list.Len()
}