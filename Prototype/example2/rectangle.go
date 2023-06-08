package main

type Rectangle struct {
	Shape
	width, height int
}

func (r *Rectangle) Clone() Cloneable {
	shape := r.Shape.Clone().(*Shape)
	return &Rectangle{*shape, r.width, r.height}
}
