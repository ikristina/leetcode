package problems

import "testing"

func TestMinStack(t *testing.T) {
	minStack := ConstructorMinStack()
	minStack.Push(1)
	minStack.Push(2)
	minStack.Push(3)
	minStack.Push(4)
	minStack.Push(5)

	if minStack.GetMin() != 1 {
		t.Errorf("Expected 1, got %d", minStack.GetMin())
	}

	minStack.Pop()
	if minStack.GetMin() != 1 {
		t.Errorf("Expected 1, got %d", minStack.GetMin())
	}

}
