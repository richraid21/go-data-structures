package structures

//Node contains the content(value) of the node and also a reference to the next Node
type Node struct {
	Value interface{}
	next *Node
}

//LinkedList contains a reference to the root Node, the head Node and the length of the List
//This is done so we have O(1) of the retrieval of the first(), peek() and len()
type LinkedList struct {
	root *Node
	head *Node
	length int
}

//Len returns the length of the LinkedList
func (l *LinkedList) Len() int{
	return l.length
}

//Append adds a Node to the end (head) of the LinkedList.
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

//RemoveAtEnd removes the last Node in the LinkedList
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

//RemoveAtStart removes the first Node in the LinkedList
//O(1)
func (l *LinkedList) RemoveAtStart() {
	if l.root == nil || l.root.next == nil {
		l.root = nil
		return
	}
	
	newRoot := l.root.next
	l.root = nil
	l.root = newRoot
	l.length = l.length -1
	return
}

//AppendList adds the provided LinkedList to the end of the referenced LinkedList
//Both LinkedLists must have the same Value Types
//O(1)
func (l *LinkedList) AppendList(lnext *LinkedList){
	l.head.next = lnext.root
	l.head = lnext.head
	l.length = l.length + lnext.length
	return
}

//Peek returns the head Node
//O(1)
func (l *LinkedList) Peek() *Node{
	return l.head
}

//Root returns the root Node
//O(1)
func (l *LinkedList) Root() *Node{
	return l.root
}

//Empty returns true if the list contains no Nodes
func (l *LinkedList) Empty() bool{
	return l.length == 0
}