package interpreter

import (
	"fmt"
	"go-funge98/eval"
	"go-funge98/eval/space"
	"go-funge98/util"
	"os"
	"regexp"

	"github.com/zeroflucs-given/generics/collections/stack"
)

type intepreterFunc func(i *Interpreter) *eval.ExitCode

type Interpreter struct {
	Ip           util.Vec
	Delta        util.Vec
	Space        *space.Space
	Stack        *stack.Stack[int]
	instructions map[rune]intepreterFunc
	InStringMode bool
}

var LINE_ENDINGS = [3]string{"\n", "\r", "\r\n"}

const MARKERS = " ;"

func MakeSpace(code string) (*space.Space, error) {
	lines := regexp.MustCompile(`\r\n|\n|\r`).Split(code, -1)

	// default to square space
	codeSpace := space.New(len(lines), len(lines), ' ')
	for x, line := range lines {
		for y, r := range line {
			switch r {
			case ' ':
				continue
			default:
				codeSpace.Set(x, y, r)
			}
		}
	}
	return &codeSpace, nil
}

func NewInterpreter(code string) (*Interpreter, error) {
	s, err := MakeSpace(code)
	if err != nil {
		return nil, err
	}
	i := Interpreter{
		Ip:           util.Vec{X: 0, Y: 0},
		Delta:        east,
		Space:        s,
		Stack:        stack.NewStack[int](16),
		InStringMode: false,
	}

	instructions := map[rune]intepreterFunc{
		// directional
		'v': (*Interpreter).PointSouth,
		'>': (*Interpreter).PointEast,
		'^': (*Interpreter).PointNorth,
		'<': (*Interpreter).PointWest,
		'?': (*Interpreter).PointRandomly,
		'[': (*Interpreter).TurnLeft,
		']': (*Interpreter).TurnRight,
		'r': (*Interpreter).Reverse,
		'x': (*Interpreter).AbsoluteVector,
		// control flow
		'#': (*Interpreter).Trampoline,
		'@': (*Interpreter).Stop,
		';': (*Interpreter).JumpOver,
		'j': (*Interpreter).JumpForward,
		'q': (*Interpreter).Quit,
		'k': (*Interpreter).Iterate,
		// decision making
		'!': (*Interpreter).LogicalNot,
		'`': (*Interpreter).GreaterThan,
		'_': (*Interpreter).EastWestIf,
		'|': (*Interpreter).NorthSouthIf,
		'w': (*Interpreter).Compare,
		// data
		// integers
		'0': (*Interpreter).PushNumber,
		'1': (*Interpreter).PushNumber,
		'2': (*Interpreter).PushNumber,
		'3': (*Interpreter).PushNumber,
		'4': (*Interpreter).PushNumber,
		'5': (*Interpreter).PushNumber,
		'6': (*Interpreter).PushNumber,
		'7': (*Interpreter).PushNumber,
		'8': (*Interpreter).PushNumber,
		'9': (*Interpreter).PushNumber,
		'a': (*Interpreter).PushNumber,
		'b': (*Interpreter).PushNumber,
		'c': (*Interpreter).PushNumber,
		'd': (*Interpreter).PushNumber,
		'e': (*Interpreter).PushNumber,
		'f': (*Interpreter).PushNumber,
		'+': (*Interpreter).Add,
		'*': (*Interpreter).Multiply,
		'-': (*Interpreter).Subtract,
		'/': (*Interpreter).Divide,
		'%': (*Interpreter).Remainder,
		// strings
		'"':  (*Interpreter).ToggleStringmode,
		'\'': (*Interpreter).FetchCharacter,
		's':  (*Interpreter).StoreCharacter,
		// stack ops
		'$':  (*Interpreter).Pop_,
		':':  (*Interpreter).Duplicate,
		'\\': (*Interpreter).Swap,
		'n':  (*Interpreter).Clear,
		// io
		'.': (*Interpreter).OutputDecimal,
		',': (*Interpreter).OutputCharacter,
		'&': (*Interpreter).InputDecimal,
		'~': (*Interpreter).InputCharacter,
	}
	i.instructions = instructions
	return &i, nil
}

func (i *Interpreter) Tick() *eval.ExitCode {
	fmt.Println(i.Ip, i.Peek(), string(i.CurrentInstruction()))
	exitCode := i.ExecuteCurrentInstruction()
	if exitCode != nil {
		return exitCode
	}
	i.Move()
	return nil
}
func (i *Interpreter) ExecuteCurrentInstruction() *eval.ExitCode {
	if i.InStringMode && i.CurrentInstruction() != '"' {
		i.Stack.Push(int(i.CurrentInstruction()))
		return nil
	}
	switch r := i.CurrentInstruction(); {
	case r == ' ':
		return nil // space is a nop
	case func() bool { _, ok := i.instructions[r]; return ok }():
		return i.instructions[r](i)
	default:
		fmt.Fprintf(os.Stderr, "unknown char `%c`\n", r)
		return nil
	}
}
func (i *Interpreter) Move() {
	i.Ip = i.Ip.Add(i.Delta)
	// TODO: implement Lahey-space wrapping
}

func (i Interpreter) CurrentInstruction() rune {
	return i.Space.GetVec(i.Ip)
}

func (i *Interpreter) Push(v int) {
	i.Stack.Push(v)
}
func (i *Interpreter) Pop() int {
	if ok, e := i.Stack.Pop(); !ok {
		return 0
	} else {
		return e
	}
}
func (i *Interpreter) Peek() int {
	if ok, e := i.Stack.Peek(); !ok {
		return 0
	} else {
		return e
	}
}
func (i Interpreter) IsEmpty() bool {
	return i.Stack.Count() == 0
}
