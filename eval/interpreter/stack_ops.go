package interpreter

import "go-funge98/eval"

func (i *Interpreter) Pop_() *eval.ExitCode {
	i.Pop()
	return nil
}
func (i *Interpreter) Duplicate() *eval.ExitCode {
	i.Push(i.Peek())
	return nil
}
func (i *Interpreter) Swap() *eval.ExitCode {
	top, next := i.Pop(), i.Pop()
	i.Push(top)
	i.Push(next)
	return nil
}
func (i *Interpreter) Clear() *eval.ExitCode {
	for !i.IsEmpty() {
		i.Pop()
	}
	return nil
}
