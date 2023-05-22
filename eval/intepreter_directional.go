package eval

import (
	"go-funge98/util"
	"math/rand"
	"time"
)

// cardinal directions
// (0,-1) (south), (1,0) (east), (0,1) (north), or (-1,0) (west)
var (
	south = util.Vec{Y: 0, X: 1}
	east  = util.Vec{Y: 1, X: 0}
	north = util.Vec{Y: 0, X: -1}
	west  = util.Vec{Y: -1, X: 0}
)

func (i *Interpreter) PointSouth() {
	i.Delta = south
}
func (i *Interpreter) PointEast() {
	i.Delta = east
}
func (i *Interpreter) PointNorth() {
	i.Delta = north
}
func (i *Interpreter) PointWest() {
	i.Delta = west
}
func (i *Interpreter) PointRandomly() {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	i.Delta = []util.Vec{south, east, north, west}[r.Intn(4)]
}
func (i *Interpreter) TurnLeft() {
	switch i.Delta {
	case south:
		i.Delta = east
	case west:
		i.Delta = south
	case north:
		i.Delta = west
	case east:
		i.Delta = north
	}
}
func (i *Interpreter) TurnRight() {
	switch i.Delta {
	case south:
		i.Delta = west
	case west:
		i.Delta = north
	case north:
		i.Delta = east
	case east:
		i.Delta = south
	}
}
func (i *Interpreter) Reverse() {
	i.Delta = i.Delta.Scale(-1)
}
