package main

import "fmt"

type Shape interface {
	Accept(v Visitor)
	Draw()
}

type Visitor interface {
	VisitDot(d *Dot)
	VisitCircle(c *Circle)
	VisitRectangle(r *Rectangle)
	VisitCompoundShape(cs *CompoundShape)
}

type Dot struct {
	ID   int
	x, y int
}

func (d *Dot) Accept(v Visitor) {
	v.VisitDot(d)
}

func (d *Dot) Draw() {
	fmt.Printf("Draw Dot, x: %d, y: %d\n", d.x, d.y)
}

type Circle struct {
	ID     int
	radius int
}

func (c *Circle) Accept(v Visitor) {
	v.VisitCircle(c)
}

func (c *Circle) Draw() {
	fmt.Printf("Draw Circle, radius: %d\n", c.radius)
}

type Rectangle struct {
	ID            int
	width, height int
}

func (r *Rectangle) Accept(v Visitor) {
	v.VisitRectangle(r)
}

func (r *Rectangle) Draw() {
	fmt.Printf("Draw Rectangle, width: %d, height: %d\n", r.width, r.height)
}

type CompoundShape struct {
	ID       int
	children []Shape
}

func (cs *CompoundShape) Accept(v Visitor) {
	v.VisitCompoundShape(cs)
}

func (cs *CompoundShape) Draw() {
	for _, shape := range cs.children {
		shape.Draw()
	}
}

type XMLExportVisitor struct{}

func (x XMLExportVisitor) VisitDot(d *Dot) {
	fmt.Printf("Exporting dot with ID %d\n", d.ID)
}

func (x XMLExportVisitor) VisitCircle(c *Circle) {
	fmt.Printf("Exporting circle with ID %d\n", c.ID)
}

func (x XMLExportVisitor) VisitRectangle(r *Rectangle) {
	fmt.Printf("Exporting rectangle with ID %d\n", r.ID)
}

func (x XMLExportVisitor) VisitCompoundShape(cs *CompoundShape) {
	fmt.Printf("Exporting compound shape with ID %d\n", cs.ID)
	for _, shape := range cs.children {
		shape.Accept(x)
	}
}

type Application struct {
	AllShapes []Shape
}

func (a Application) Export() {
	exportVisitor := XMLExportVisitor{}

	for _, shape := range a.AllShapes {
		shape.Accept(exportVisitor)
	}
}

func main() {
	dot := &Dot{ID: 1, x: 2, y: 3}
	circle := &Circle{ID: 2, radius: 3}
	rectangle := &Rectangle{ID: 3, width: 4, height: 5}
	compoundshape := &CompoundShape{ID: 4, children: []Shape{dot, circle, rectangle}}
	app := Application{
		AllShapes: []Shape{
			dot,
			circle,
			rectangle,
			compoundshape,
		},
	}

	app.Export()
}
