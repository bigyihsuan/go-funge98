package stackstack

import "github.com/zeroflucs-given/generics/collections/stack"

type StackStack[T any] struct {
	stack *stack.Stack[*stack.Stack[T]]
}

func New[T any]() StackStack[T] {
	var ss StackStack[T]
	ss.stack = stack.NewStack[*stack.Stack[T]](16)
	inner := stack.NewStack[T](16)
	ss.stack.Push(inner)
	return ss
}

func (ss *StackStack[T]) PushCell(cell T) {
	if ok, top := ss.stack.Peek(); ok {
		top.Push(cell)
	}
}
func (ss *StackStack[T]) PopCell() T {
	var v T
	if ok, top := ss.stack.Peek(); !ok {
		return v
	} else if ok, cell := top.Pop(); !ok {
		return v
	} else {
		return cell
	}
}
func (ss *StackStack[T]) PeekCell() T {
	var v T
	if ok, top := ss.stack.Peek(); !ok {
		return v
	} else if ok, cell := top.Peek(); !ok {
		return v
	} else {
		return cell
	}
}

func (ss *StackStack[T]) PushStack(s *stack.Stack[T]) {
	ss.stack.Push(s)
}
func (ss *StackStack[T]) PopStack() (bool, *stack.Stack[T]) {
	return ss.stack.Pop()
}
