package interpreter

import (
	"go-funge98/eval"
)

func (i *Interpreter) ToggleStringmode() *eval.ExitCode {
	i.InStringMode = !i.InStringMode
	return nil
}
func (i *Interpreter) FetchCharacter() *eval.ExitCode {
	i.Move()
	r := i.CurrentInstruction()
	i.Push(int(r))
	return nil
}
func (i *Interpreter) StoreCharacter() *eval.ExitCode {
	r := rune(i.Pop())
	ip := i.Ip
	i.Move()
	i.Space.SetVec(i.Ip, r)
	i.Ip = ip
	return nil
}
