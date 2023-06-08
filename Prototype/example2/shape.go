package main

type Cloneable interface {
	Clone() Cloneable
}

type Shape struct {
	x, y  int
	color string
}

func (s *Shape) Clone() Cloneable {
	return &Shape{s.x, s.y, s.color}
}
