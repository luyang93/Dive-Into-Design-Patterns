package main

import "fmt"

type Shape interface {
	Draw(x, y, radius int)
}

type CircleFlyweight struct {
	color string
}

func NewCircleFlyweight(color string) *CircleFlyweight {
	return &CircleFlyweight{color: color}
}

func (c *CircleFlyweight) Draw(x, y, radius int) {
	fmt.Printf("Drawing circle of color %s at position (%d, %d) with radius %d\n", c.color, x, y, radius)
}

type ShapeFactory struct {
	circleMap map[string]*CircleFlyweight
}

func NewShapeFactory() *ShapeFactory {
	return &ShapeFactory{circleMap: make(map[string]*CircleFlyweight)}
}

func (sf *ShapeFactory) GetCircle(color string) *CircleFlyweight {
	if circle, found := sf.circleMap[color]; found {
		return circle
	} else {
		circle = NewCircleFlyweight(color)
		sf.circleMap[color] = circle
		fmt.Printf("Creating circle of color: %s\n", color)
		return circle
	}
}

func main() {
	sf := NewShapeFactory()

	redCircle := sf.GetCircle("red")
	redCircle.Draw(10, 10, 50)

	blueCircle := sf.GetCircle("blue")
	blueCircle.Draw(20, 20, 60)

	anotherRedCircle := sf.GetCircle("red")
	anotherRedCircle.Draw(30, 30, 40)
}
