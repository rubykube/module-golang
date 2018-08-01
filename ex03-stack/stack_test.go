package stack

import "testing"

func TestStack(t *testing.T) {
	var stack *Stack = New()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	for i := 5; i > 0; i-- {
		item := stack.Pop()

		if item != i {
			t.Error("TestStack failed...", i)
		}
	}

	stack.Push(11)
	stack.Push(12)
	stack.Push(13)
	stack.Push(14)
	stack.Push(15)

	for i := 15; i > 10; i-- {
		item := stack.Pop()

		if item != i {
			t.Error("TestStack failed...", i)
		}
	}
}
