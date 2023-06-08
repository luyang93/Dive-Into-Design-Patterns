package main

type Circle struct {
	Shape
	radius int
}

func (c *Circle) Clone() Cloneable {
	shape := c.Shape.Clone().(*Shape)
	return &Circle{*shape, c.radius}
}
