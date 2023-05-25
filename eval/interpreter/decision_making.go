package interpreter

import "go-funge98/eval"

func (i *Interpreter) LogicalNot() *eval.ExitCode {
	if i.Pop() == 0 {
		i.Push(1)
	} else {
		i.Push(0)
	}
	return nil
}
func (i *Interpreter) GreaterThan() *eval.ExitCode {
	r := i.Pop()
	l := i.Pop()
	if l > r {
		i.Push(1)
	} else {
		i.Push(0)
	}
	return nil
}
func (i *Interpreter) EastWestIf() *eval.ExitCode {
	if i.Pop() == 0 {
		i.PointEast()
	} else {
		i.PointWest()
	}
	return nil
}
func (i *Interpreter) NorthSouthIf() *eval.ExitCode {
	if i.Pop() == 0 {
		i.PointNorth()
	} else {
		i.PointSouth()
	}
	return nil
}
func (i *Interpreter) Compare() *eval.ExitCode {
	r := i.Pop()
	l := i.Pop()
	if l < r {
		i.TurnLeft()
	} else if l > r {
		i.TurnRight()
	} // else nop
	return nil
}
