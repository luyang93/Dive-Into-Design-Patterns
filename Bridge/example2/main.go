package main

import "fmt"

type Painter interface {
	Paint(*Shape)
}

type RedPainter struct{}

func (rp *RedPainter) Paint(s *Shape) {
	s.color = "Red"
	fmt.Println("Shape: " + s.shape + ", Color: " + s.color)
}

type BluePainter struct{}

func (bp *BluePainter) Paint(s *Shape) {
	s.color = "Blue"
	fmt.Println("Shape: " + s.shape + ", Color: " + s.color)
}

type Shape struct {
	shape   string
	color   string
	painter Painter
}

type Circle struct {
	Shape
}

func NewCircle() *Circle {
	return &Circle{
		Shape{
			shape: "Circle",
		},
	}
}

type Square struct {
	Shape
}

func NewSquare() *Square {
	return &Square{
		Shape{
			shape: "Square",
		},
	}
}

func (s *Shape) Draw() {
	fmt.Println("Drawing " + s.color + " " + s.shape)
	s.painter.Paint(s)
}

func (s *Shape) SetPainter(painter Painter) {
	s.painter = painter
}

func main() {
	redPainter := &RedPainter{}
	bluePainter := &BluePainter{}

	circle := NewCircle()
	circle.SetPainter(redPainter)
	circle.Draw()

	circle.SetPainter(bluePainter)
	circle.Draw()

	square := NewSquare()
	square.SetPainter(redPainter)
	square.Draw()

	square.SetPainter(bluePainter)
	square.Draw()
}
