package structures

import "testing"

func TestQueueCreation(t *testing.T){
	t.Log("Test the creation of a new Queue.")
	q := NewQueue()
	
	if q == nil {
		t.Error("NewQueue() returned nil")
	}
	
	if q.Len() != 0 {
		t.Error("New Queue should have length of 0")
	}
	
	if q.Empty() != true {
		t.Error("New Queue should be empty.")
	}
	
}

func TestQueueEnqueue(t *testing.T){
	t.Log("Test the addition of elements to a Queue.")
	q := NewQueue()
	
	node1 := Node{Value: "First Element"}
	q.Enqueue(&node1)
	
	if q.Len() != 1 {
		t.Error("Queue should have a length of 1 after a single element enqueued.")
	}
	
	if q.Peek() != &node1 {
		t.Error("Enqueued element should match Peek() operation for the first element.")
	}
	
	if q.Empty() == true {
		t.Error("Queue should not return Empty() == true after an enqueue operation.")
	}
	
	node2 := Node{Value: "Second Element"}
	q.Enqueue(&node2)
	
	if q.Len() != 2 {
		t.Error("Queue should have a length of 2 after a second element enqueued.")
	}
	
	if q.Peek() != &node1 {
		t.Error("Peek() operation should point to the oldest element.")
	}
	
}

func TestQueueDequeue(t *testing.T){
	t.Log("Test the removal of elements to a Queue.")
	q := NewQueue()
	
	node1 := Node{Value: "First Element"}
	q.Enqueue(&node1)
	
	node2 := Node{Value: "Second Element"}
	q.Enqueue(&node2)
	
	node3 := Node{Value: "Third Element"}
	q.Enqueue(&node3)
	
	if q.Len() != 3 {
		t.Error("Queue should contain 3 elements after 3 enqueue operations.")
	}
	
	if q.Dequeue() != &node1 {
		t.Error("Dequeue should return the oldest element in the queue.")
	}
	
	if q.Len() != 2 {
		t.Error("Queue length should decrement following a Dequeue() operation.")
	}
	
	if q.Peek() != &node2 {
		t.Error("Peek() operation should refer to the 2nd oldest element after a Dequeue() of the oldest.")
	}
	
}