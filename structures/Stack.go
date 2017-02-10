package structures

type stack struct {
	list *LinkedList
}

func (s *stack) Push(i int) {
	n := Node{data: i}
	s.list.Append(&n)
}

func NewStack() *stack {
	l := new(LinkedList)
	s := stack{list: l}
	return &s
}

func (s *stack) Pop() int {
	i := s.Peek()
	s.list.RemoveAtEnd()
	return i
}

func (s *stack) Peek() int {
	return s.list.Peek().data
}

func (s *stack) Empty() bool {
	return s.list.Empty()
}

func (s *stack) Depth() int {
	return s.list.Len()
}