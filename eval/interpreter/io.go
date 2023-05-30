package interpreter

import (
	"fmt"
	"go-funge98/eval"
	"os"
)

func (i *Interpreter) OutputDecimal() *eval.ExitCode {
	c := i.Pop()
	_, err := fmt.Print(rune(c))
	if err != nil {
		i.Reverse()
	}
	return nil
}
func (i *Interpreter) OutputCharacter() *eval.ExitCode {
	c := i.Pop()
	_, err := fmt.Print(string(rune(c)))
	if err != nil {
		i.Reverse()
	}
	return nil
}
func (i *Interpreter) InputDecimal() *eval.ExitCode {
	var n int
	fmt.Fscanf(os.Stdin, "%d", &n)
	i.Push(n)
	return nil
}
func (i *Interpreter) InputCharacter() *eval.ExitCode {
	var n rune
	fmt.Fscanf(os.Stdin, "%c", &n)
	i.Push(int(n))
	return nil
}
