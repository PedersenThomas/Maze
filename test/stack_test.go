package test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/pedersenthomas/Maze/utils"
)

func TestPushCorrectCount(t *testing.T) {
	stack, num := createStack()
	if stack.Count() != num {
		t.Errorf("Push() %d items, got %d", num, stack.Count())
	}
}

func TestPopEmptyStack(t *testing.T) {
	stack := utils.Stack{}
	item := stack.Pop()
	if item != nil {
		t.Errorf("Pop() from empty stack, got %v", item)
	}
}

func TestPushTwicePopTwiceEmpty(t *testing.T) {
	stack := utils.Stack{}
	stack.Push(1)
	stack.Push(2)
	item := stack.Pop()
	if item != 2 {
		t.Errorf("Pop(), should get 2, got %v", item)
	}

	item2 := stack.Pop()
	if item2 != 1 {
		t.Errorf("Pop(), should get 1, got %v", item2)
	}

	if stack.Count() != 0 {
		t.Errorf("Pop(), count should be 0, go %d", stack.Count())
	}
}

func TestIsEmpty(t *testing.T) {
	stack, _ := createStack()
	for stack.Count() > 0 {
		stack.Pop()
	}

	result := stack.IsEmpty()
	if !result {
		t.Errorf("IsEmpty() should be true, got %v", result)
	}
}

func TestPeekNoCountChange(t *testing.T) {
	stack, num := createStack()
	stack.Peek()
	if stack.Count() != num {
		t.Errorf("Peek() should not remove item. Had %d, now %d", num, stack.Count())
	}
}

func TestFixedSizeStackPeekNoChange(t *testing.T) {
	stack := utils.Stack{}
	stack.Push(1)
	stack.Push(2)
	item := stack.Peek()
	if item != 2 {
		t.Errorf("Pop(), should get 2, got %v", item)
	}
}

func BenchmarkPush(b *testing.B) {
	stack := utils.Stack{}
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}

func createStack() (utils.Stack, int) {
	rand.Seed(time.Now().UTC().UnixNano())
	num := rand.Intn(100) + 1
	stack := utils.Stack{}
	for i := 0; i < num; i++ {
		stack.Push(i)
	}
	return stack, num
}
