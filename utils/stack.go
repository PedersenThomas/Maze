package utils

type Stack struct {
	count int
	items []interface{}
}

func (stack *Stack) Count() int {
	return len(stack.items)
}

func (stack *Stack) IsEmpty() bool {
	return stack.Count() == 0
}

func (stack *Stack) Pop() interface{} {
	if stack.Count() == 0 {
		return nil
	}

	item := stack.items[stack.Count()-1]
	stack.items = stack.items[:stack.Count()-1]
	return item
}

func (stack *Stack) Push(item interface{}) {
	stack.items = append(stack.items, item)
	stack.count += 1
}

func (stack *Stack) Peek() interface{} {
	if stack.Count() == 0 {
		return nil
	}

	return stack.items[stack.Count()-1]
}
