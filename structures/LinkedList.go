package structures

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	root *Node
	head *Node
	length int
}

func (l *LinkedList) Len() int{
	return l.length
}

//O(1)
func (l *LinkedList) Append(n *Node){
	l.length = l.length + 1
	
	if l.root == nil {
		l.root = n
		l.head = n
		return
	}
	l.head.next = n
	l.head = n
	
	return
}

//O(n)
func (l *LinkedList) RemoveAtEnd(){
	newHead := l.root
	
	for newHead.next != l.head{
		newHead = newHead.next
	}
	
	newHead.next = nil
	l.head = newHead
	l.length = l.length - 1
	
}

//O(1)
func (l *LinkedList) AppendList(lnext *LinkedList){
	l.head.next = lnext.root
	l.head = lnext.head
	l.length = l.length + lnext.length
	return
}

//O(1)
func (l *LinkedList) Peek() *Node{
	return l.head
}

//O(1)
func (l *LinkedList) Empty() bool{
	return l.length == 0
}