package main

import (
	"fmt"
	"go-funge98/eval"
	"go-funge98/eval/interpreter"
	"os"

	"github.com/jessevdk/go-flags"
)

func main() {
	var opts struct {
		File  flags.Filename `short:"f" long:"file" value-name:"FILE" description:"Input code file." required:"true"`
		Debug bool           `short:"d" long:"debug" description:"Output debugging information."`
	}
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("file:%v\n", opts.File)
	fmt.Printf("debug:%v\n", opts.Debug)

	interpreter, err := interpreter.NewInterpreter(string(opts.File))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(interpreter.Space)
	var exitCode *eval.ExitCode
	for {
		exitCode = interpreter.Tick()
		if exitCode != nil {
			break
		}
	}

	_, s := interpreter.Stack.PopStack()
	for s.Count() > 0 {
		_, e := s.Pop()
		fmt.Printf("%v ", e)
	}
	fmt.Println()

	if exitCode != nil {
		os.Exit(exitCode.Code)
	}
}
