package datastructure

import "testing"

func TestStackPush(t *testing.T) {
	stack := Stack{}
	stack.Push("hello")
	if stack.IsEmpty() {
		t.Errorf("test failed since the stack still empty")
	}
}

func TestStackPop(t *testing.T) {
	stack := Stack{
		"hello",
	}
	got, isEmpty := stack.Pop()
	if (!isEmpty) {
		t.Errorf("test failed since the stack was empty")
	}
	if (got != "hello") {
		t.Errorf("test failed since return value is not hello")
	}
}