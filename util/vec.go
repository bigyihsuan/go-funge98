package util

type Vec struct {
	X int
	Y int
}

func (v Vec) Add(o Vec) Vec {
	return Vec{
		X: v.X + o.X,
		Y: v.Y + o.Y,
	}
}

func (v Vec) Sub(o Vec) Vec {
	return Vec{
		X: v.X - o.X,
		Y: v.Y - o.Y,
	}
}

func (v Vec) Scale(c int) Vec {
	return Vec{
		X: c * v.X,
		Y: c * v.Y,
	}
}
