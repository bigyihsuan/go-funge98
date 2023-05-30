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
		File  flags.Filename `short:"f" long:"file" value-name:"FILE" description:"Input code file."`
		Code  flags.Filename `short:"c" long:"code" value-name:"CODE" description:"Argument-provided Befunge code."`
		Debug bool           `short:"d" long:"debug" description:"Output debugging information."`
	}
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	if opts.File != "" && opts.Code != "" {
		fmt.Fprintln(os.Stderr, "-f/--file and -c/--code flags are mutually exclusive")
		os.Exit(1)
	}

	var code string

	if opts.File != "" {
		bytes, err := os.ReadFile(string(opts.File))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		code = string(bytes)
	} else if opts.Code != "" {
		code = string(opts.Code)
	}

	// fmt.Printf("file:%v\n", opts.File)
	// fmt.Printf("debug:%v\n", opts.Debug)

	interpreter, err := interpreter.NewInterpreter(code)
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

	for interpreter.Stack.Count() > 0 {
		_, e := interpreter.Stack.Pop()
		fmt.Printf("%v ", e)
	}
	fmt.Println()

	if exitCode != nil {
		os.Exit(exitCode.Code)
	}
}
