package main

import (
	"fmt"
	"math"
)

type Shape interface {
	GetArea() float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) GetArea() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Square struct {
	Side float64
}

func (s *Square) GetArea() float64 {
	return s.Side * s.Side
}

type ShapeGroup struct {
	Shapes []Shape
}

func (sg *ShapeGroup) AddShape(shape Shape) {
	sg.Shapes = append(sg.Shapes, shape)
}

func (sg *ShapeGroup) GetArea() float64 {
	totalArea := 0.0
	for _, shape := range sg.Shapes {
		totalArea += shape.GetArea()
	}
	return totalArea
}

func main() {
	circle := Circle{Radius: 5}
	square := Square{Side: 4}

	shapeGroup1 := ShapeGroup{}
	shapeGroup1.AddShape(&circle)
	shapeGroup1.AddShape(&square)

	circle2 := Circle{Radius: 3}
	shapeGroup2 := &ShapeGroup{}
	shapeGroup2.AddShape(&circle2)

	mainGroup := ShapeGroup{}
	mainGroup.AddShape(&shapeGroup1)
	//mainGroup.AddShape(&shapeGroup2)

	fmt.Println("Total area:", mainGroup.GetArea())
}
