package structures

import "testing"

func TestListLength(t *testing.T) {
	line := new(LinkedList)
	n := Node{data: 5}
	line.Append(&n)
	
	if line.Len() != 1 {
		t.Error("1 Element List Should have length of 1")
	}
}

func TestListRemoveAtEnd(t *testing.T) {
	line:= new(LinkedList)
	
	first := Node{data: 8}
	line.Append(&first)
	
	second := Node{data: 2}
	line.Append(&second)
	
	line.RemoveAtEnd()
	
	if line.Peek() != &first || line.Len() != 1{
		t.Error("Removal of End Element Incorrect")
	}
}

func TestListRemoveAtStart(t *testing.T) {
	line:= new(LinkedList)
	
	first := Node{data: 8}
	line.Append(&first)
	
	second := Node{data: 2}
	line.Append(&second)
	
	line.RemoveAtStart()
	
	if line.Peek() != &second || line.Len() != 1{
		t.Error("Removal of Start Element Incorrect")
	}
}

func TestListAppendList(t *testing.T) {
	list1 := new(LinkedList)
	list2 := new(LinkedList)
	
	for i:=0;i<5;i++ {
		n := Node{data: i}
		list1.Append(&n)
	}
	
	for i:=5;i<10;i++ {
		n := Node{data: i}
		list2.Append(&n)
	}
	
	list1.AppendList(list2)
	
	if list1.Len() != 10 || list1.Peek().data != 9 {
		t.Error("List Append List")
	}
	
}

func TestEmptyList(t *testing.T) {
	list1 := new(LinkedList)
	if !list1.Empty(){
		t.Error("New List Should be empty")
	}
}