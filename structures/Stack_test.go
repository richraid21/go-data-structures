package structures

import "testing"

func TestStackCreation(t *testing.T) {
	t.Log("Test the creation of a Stack.")
	s := NewStack()
	
	if s == nil {
		t.Error("Stack should have been created")
	}
}

func TestStackPush(t *testing.T) {
	t.Log("Test the push operation on an existing Stack.")
	s := NewStack()
	
	s.Push(&Node{Value: 17})
	s.Push(&Node{Value: 72})
	
	if s.Depth() != 2 {
		t.Error("Stack Push 3 Elements should have depth of 3")
	}
}

func TestStackPeek(t *testing.T) {
	t.Log("Test the peek operation on an existing Stack.")
	s := NewStack()
	
	s.Push(&Node{Value: 17})
	s.Push(&Node{Value: 72})
	
	if s.Peek().Value != 72 {
		t.Error("Stack Peek() should return last pushed element")
	}
	
	if s.Depth() != 2 {
		t.Error("Stack Peek() should not affect depth of Stack")
	}
}

func TestStackPop(t *testing.T) {
	t.Log("Test the pop operation on an existing Stack.")
	s := NewStack()
	
	s.Push(&Node{Value: 17})
	s.Push(&Node{Value: 72})
	
	if s.Pop().Value != 72 {
		t.Error("Stack Pop() should return last pushed element")
	}
	
	if s.Depth() != 1 {
		t.Error("Stack Pop() should remove 1 element")
	}
}