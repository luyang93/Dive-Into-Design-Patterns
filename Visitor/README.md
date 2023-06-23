```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface Shape {
    +Accept(v: Visitor)
    +Draw()
}

interface Visitor {
    +VisitDot(d: Dot)
    +VisitCircle(c: Circle)
    +VisitRectangle(r: Rectangle)
    +VisitCompoundShape(cs: CompoundShape)
}

class Dot {
    -ID: int
    -x: int
    -y: int
    +Accept(v: Visitor)
    +Draw()
}

class Circle {
    -ID: int
    -radius: int
    +Accept(v: Visitor)
    +Draw()
}

class Rectangle {
    -ID: int
    -width: int
    -height: int
    +Accept(v: Visitor)
    +Draw()
}

class CompoundShape {
    -ID: int
    -children: []Shape
    +Accept(v: Visitor)
    +Draw()
}

class XMLExportVisitor {
    +VisitDot(d: Dot)
    +VisitCircle(c: Circle)
    +VisitRectangle(r: Rectangle)
    +VisitCompoundShape(cs: CompoundShape)
}

class Application {
    -AllShapes: []Shape
    +Export()
}

Dot --> Shape: implements
Circle --> Shape: implements
Rectangle --> Shape: implements
CompoundShape --> Shape: implements

XMLExportVisitor --> Visitor: implements

Shape --> Visitor: uses
CompoundShape --> Shape: uses
Application --> Shape: uses
Application --> XMLExportVisitor: uses
```