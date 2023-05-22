package eval

import (
	"go-funge98/eval/space"
	"os"
	"regexp"
)

type Intepreter struct {
	instructionPointer Point
	spacePointer       Point
	space              *space.Space
}

var lineEndings = [3]string{"\n", "\r", "\r\n"}

func LoadFile(filename string) (*space.Space, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	code := string(bytes)
	lines := regexp.MustCompile("\r\n|\n|\r").Split(code, -1)

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

func NewInterpreter(filename string) (*Intepreter, error) {
	s, err := LoadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Intepreter{
		instructionPointer: Point{X: 0, Y: 0},
		spacePointer:       Point{X: 0, Y: 0},
		space:              s,
	}, nil
}

type Point struct {
	X int
	Y int
}
