package main

import "fmt"

type Graphic interface {
	Move(x, y int)
	Draw()
}

func Equal(g1, g2 Graphic) bool {
	switch g1 := g1.(type) {
	case *Dot:
		if g2, ok := g2.(*Dot); ok {
			return g1.x == g2.x && g1.y == g2.y
		}
	case *Circle:
		if g2, ok := g2.(*Circle); ok {
			return g1.x == g2.x && g1.y == g2.y && g1.radius == g2.radius
		}
	}
	return false
}

type Dot struct {
	x, y int
}

func (d *Dot) Move(x, y int) {
	d.x += x
	d.y += y
}

func (d *Dot) Draw() {
	fmt.Printf("Drawing a dot at (%d, %d)\n", d.x, d.y)
}

func NewDot(x, y int) *Dot {
	return &Dot{x, y}
}

type Circle struct {
	Dot
	radius int
}

func (c *Circle) Draw() {
	fmt.Printf("Drawing a circle at (%d, %d) with radius %d\n", c.x, c.y, c.radius)
}

func NewCircle(x, y, radius int) *Circle {
	return &Circle{Dot{x, y}, radius}
}

type CompoundGraphic struct {
	children []Graphic
}

func (cg *CompoundGraphic) Add(child Graphic) {
	cg.children = append(cg.children, child)
}

func (cg *CompoundGraphic) Remove(child Graphic) {
	var newChildren []Graphic
	for _, c := range cg.children {
		if !Equal(c, child) {
			newChildren = append(newChildren, c)
		}
	}
	cg.children = newChildren
}

func (cg *CompoundGraphic) Move(x, y int) {
	for _, child := range cg.children {
		child.Move(x, y)
	}
}

func (cg *CompoundGraphic) Draw() {
	for _, child := range cg.children {
		child.Draw()
	}
}

type ImageEditor struct {
	compoundGraphic *CompoundGraphic
}

func (ie *ImageEditor) Load() {
	ie.compoundGraphic = &CompoundGraphic{}
	ie.compoundGraphic.Add(NewDot(1, 2))
	ie.compoundGraphic.Add(NewDot(2, 3))
	ie.compoundGraphic.Add(NewCircle(5, 3, 10))
	ie.compoundGraphic.Add(NewCircle(6, 4, 11))

}

func (ie *ImageEditor) GroupSelected(components []Graphic) {
	group := &CompoundGraphic{}
	for _, component := range components {
		group.Add(component)
		ie.compoundGraphic.Remove(component)
	}
	ie.compoundGraphic.Add(group)
	ie.compoundGraphic.Draw()
}

func main() {
	imageEditor := &ImageEditor{}
	imageEditor.Load()
	fmt.Println(imageEditor.compoundGraphic)
	fmt.Println(imageEditor.compoundGraphic.children)

	imageEditor.GroupSelected([]Graphic{NewDot(2, 3), NewDot(1, 2), NewDot(3, 4)})

	fmt.Println(imageEditor.compoundGraphic.children)
	fmt.Println(imageEditor.compoundGraphic.children[2])
	fmt.Println(imageEditor.compoundGraphic.children[1].(*Circle))
	fmt.Println(imageEditor.compoundGraphic.children[2].(*CompoundGraphic).children[0])
}
