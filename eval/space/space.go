package space

import (
	"go-funge98/util"
	"strings"
)

// the Funge-Space
type Space struct {
	space  [][]rune
	filler rune
}

// Makes a new Space with the given size, and with the given filler cell (for "empty" cells).
func New(xSize, ySize int, filler rune) Space {
	var s Space
	s.space = make([][]rune, xSize)
	for i := range s.space {
		s.space[i] = make([]rune, ySize)
		for j := range s.space[i] {
			s.space[i][j] = filler
		}
	}
	s.filler = filler
	return s
}

// the horizontal size
func (s Space) Xsize() int {
	return len(s.space)
}

// the vertical size; returns -1 if the xsize is empty
func (s Space) Ysize() int {
	if s.Xsize() == 0 {
		return -1
	}
	return len(s.space[0])
}

// Resize the space to fit the given indexes
func (s *Space) resize(x, y int) {
	// don't resize if there the indexes are inside the space
	if x < s.Xsize() && y < s.Ysize() {
		return
	}
	newX, newY := s.Xsize(), s.Ysize()
	if x >= newX {
		newX = x + 1
	}
	if y >= newY {
		newY = y + 1
	}
	newSpace := New(newX, newY, s.filler)
	for i := range s.space {
		copy(newSpace.space[i], s.space[i])
	}
	s.space = newSpace.space
}

// Get the cell at this (x,y) location.
// x and y can be negative. 0-indexed.
func (s Space) Get(x, y int) (t rune) {
	s.resize(x, y)
	return s.space[x][y]
}

func (s Space) GetVec(v util.Vec) (t rune) {
	return s.Get(v.X, v.Y)
}

// Set the cell at this (x,y) location.
// x and y can be negative. 0-indexed.
func (s *Space) Set(x, y int, v rune) {
	s.resize(x, y)
	s.space[x][y] = v
}

// fmt.Stringer
func (s Space) String() string {
	var builder strings.Builder
	for i, col := range s.space {
		for _, cell := range col {
			builder.WriteRune(cell)
		}
		if i < s.Xsize()-1 {
			builder.WriteString("\n")
		}
	}
	return builder.String()
}
