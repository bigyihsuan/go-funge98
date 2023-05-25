package interpreter

import (
	"go-funge98/eval"
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

func (i *Interpreter) PointSouth() (exit *eval.ExitCode) {
	i.Delta = south
	return
}
func (i *Interpreter) PointEast() (exit *eval.ExitCode) {
	i.Delta = east
	return
}
func (i *Interpreter) PointNorth() (exit *eval.ExitCode) {
	i.Delta = north
	return
}
func (i *Interpreter) PointWest() (exit *eval.ExitCode) {
	i.Delta = west
	return
}
func (i *Interpreter) PointRandomly() (exit *eval.ExitCode) {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	i.Delta = []util.Vec{south, east, north, west}[r.Intn(4)]
	return
}
func (i *Interpreter) TurnLeft() (exit *eval.ExitCode) {
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
	return
}
func (i *Interpreter) TurnRight() (exit *eval.ExitCode) {
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
	return
}
func (i *Interpreter) Reverse() (exit *eval.ExitCode) {
	i.Delta = i.Delta.Scale(-1)
	return
}
func (i *Interpreter) AbsoluteVector() (exit *eval.ExitCode) {
	dy := i.Pop()
	dx := i.Pop()
	i.Delta = util.Vec{X: dx, Y: dy}
	return
}
