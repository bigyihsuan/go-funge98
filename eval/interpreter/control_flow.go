package interpreter

import (
	"go-funge98/eval"
	"strings"
)

func (i *Interpreter) Trampoline() (exit *eval.ExitCode) {
	i.Move()
	return
}
func (i *Interpreter) Stop() (exit *eval.ExitCode) {
	return &eval.ExitCode{Code: 0}
}
func (i *Interpreter) JumpOver() (exit *eval.ExitCode) {
	i.Move() // skip current ;
	for i.CurrentInstruction() != ';' {
		i.Move()
	}
	return
}
func (i *Interpreter) JumpForward() (exit *eval.ExitCode) {
	steps := i.Pop()
	lastDelta := i.Delta
	if steps < 0 {
		i.Reverse()
		steps = -steps
	}
	for s := 0; s < steps; s++ {
		i.Move()
	}
	i.Delta = lastDelta
	return
}
func (i *Interpreter) Quit() (exit *eval.ExitCode) {
	return &eval.ExitCode{Code: i.Pop()}
}
func (i *Interpreter) Iterate() (exit *eval.ExitCode) {
	// find next instruction
	currentIp := i.Ip
	i.Move()
	for strings.ContainsRune(MARKERS, i.CurrentInstruction()) { // move to the next non-marker instruction
		i.Move()
	}
	count := i.Pop()
	for n := 0; n < count; n++ {
		i.ExecuteCurrentInstruction()
	}
	i.Ip = currentIp
	return
}
