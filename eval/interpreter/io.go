package interpreter

import (
	"fmt"
	"go-funge98/eval"
)

func (i *Interpreter) Print() *eval.ExitCode {
	c := i.Stack.PopCell()
	fmt.Print(string(rune(c)))
	return nil
}
