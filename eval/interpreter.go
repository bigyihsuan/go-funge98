package eval

import (
	"fmt"
	"go-funge98/eval/space"
	"go-funge98/eval/stackstack"
	"go-funge98/util"
	"os"
	"regexp"
)

type Interpreter struct {
	Ip    util.Vec
	Delta util.Vec
	Space *space.Space
	Stack stackstack.StackStack[int]
}

var lineEndings = [3]string{"\n", "\r", "\r\n"}

func LoadFile(filename string) (*space.Space, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	code := string(bytes)
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

func NewInterpreter(filename string) (*Interpreter, error) {
	s, err := LoadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Interpreter{
		Ip:    util.Vec{X: 0, Y: 0},
		Delta: util.Vec{X: 0, Y: 0},
		Space: s,
		Stack: stackstack.New[int](),
	}, nil
}

func (i *Interpreter) Tick() {
	i.Exec()
	fmt.Printf("curr: `%c`\n", i.Space.GetVec(i.Ip))
	i.Move()
}
func (i *Interpreter) Exec() {
	switch r := i.Space.GetVec(i.Ip); r {
	case ' ':
		return // space is a nop
	// directional instructions
	case 'v':
		i.PointSouth()
	case '>':
		i.PointEast()
	case '^':
		i.PointNorth()
	case '<':
		i.PointWest()
	case '?':
		i.PointRandomly()
	case '[':
		i.TurnLeft()
	case ']':
		i.TurnRight()
	case 'r':
		i.Reverse()
	case 'x':
		/*
			TODO: pop a vector off the stack, and set delta to that vector
			a "vector" on the stack is (dy dx) (dy on top)
		*/
	default:
		fmt.Fprintf(os.Stderr, "unknown char `%c`\n", r)
		return
	}
}
func (i *Interpreter) Move() {
	i.Ip = i.Ip.Add(i.Delta)
	// TODO: implement Lahey-space wrapping
}
