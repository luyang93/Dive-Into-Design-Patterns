```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

class ShapeFactory {
  - circleMap : map[string]*CircleFlyweight
  + GetCircle(color : string) : *CircleFlyweight
}

class CircleFlyweight {
  - color : string
  + Draw(x, y, radius : int)
}

interface Shape {
  + Draw(x, y, radius : int)
}

ShapeFactory --> CircleFlyweight : creates >
Shape <|-- CircleFlyweight
```