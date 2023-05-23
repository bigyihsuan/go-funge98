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
	i.Stack.PushCell(int(n))
	return nil
}
func (i *Interpreter) Add() (exit *eval.ExitCode) {
	r := i.Stack.PopCell()
	l := i.Stack.PopCell()
	i.Stack.PushCell(l + r)
	return nil
}
func (i *Interpreter) Multiply() (exit *eval.ExitCode) {
	r := i.Stack.PopCell()
	l := i.Stack.PopCell()
	i.Stack.PushCell(l * r)
	return nil
}
func (i *Interpreter) Subtract() (exit *eval.ExitCode) {
	r := i.Stack.PopCell()
	l := i.Stack.PopCell()
	i.Stack.PushCell(l - r)
	return nil
}
func (i *Interpreter) Divide() (exit *eval.ExitCode) {
	r := i.Stack.PopCell()
	l := i.Stack.PopCell()
	if r == 0 {
		i.Stack.PushCell(0)
	} else {
		i.Stack.PushCell(l / r)
	}
	return nil
}
func (i *Interpreter) Remainder() (exit *eval.ExitCode) {
	r := i.Stack.PopCell()
	l := i.Stack.PopCell()
	if r == 0 {
		i.Stack.PushCell(0)
	} else {
		i.Stack.PushCell(l % r)
	}
	return nil
}
