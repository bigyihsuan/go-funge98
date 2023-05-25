package interpreter

import (
	"go-funge98/eval"
	"strconv"
)

func (i *Interpreter) PushNumber() (exit *eval.ExitCode) {
	n, err := strconv.ParseInt(string(i.CurrentInstruction()), 16, 64)
	if err != nil {
		return
	}
	i.Push(int(n))
	return nil
}
func (i *Interpreter) Add() (exit *eval.ExitCode) {
	r := i.Pop()
	l := i.Pop()
	i.Push(l + r)
	return nil
}
func (i *Interpreter) Multiply() (exit *eval.ExitCode) {
	r := i.Pop()
	l := i.Pop()
	i.Push(l * r)
	return nil
}
func (i *Interpreter) Subtract() (exit *eval.ExitCode) {
	r := i.Pop()
	l := i.Pop()
	i.Push(l - r)
	return nil
}
func (i *Interpreter) Divide() (exit *eval.ExitCode) {
	r := i.Pop()
	l := i.Pop()
	if r == 0 {
		i.Push(0)
	} else {
		i.Push(l / r)
	}
	return nil
}
func (i *Interpreter) Remainder() (exit *eval.ExitCode) {
	r := i.Pop()
	l := i.Pop()
	if r == 0 {
		i.Push(0)
	} else {
		i.Push(l % r)
	}
	return nil
}
