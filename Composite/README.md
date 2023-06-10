```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface Graphic {
    +Move(x, y: int)
    +Draw()
    +Equal(g: Graphic): bool
}

class Dot {
    -x, y: int
    +Move(x, y: int)
    +Draw()
    +Equal(g: Graphic): bool
}

class Circle {
    -Dot
    -radius: int
    +Draw()
    +Equal(g: Graphic): bool
}

class CompoundGraphic {
    -children: []Graphic
    +Add(child: Graphic)
    +Remove(child: Graphic)
    +Move(x, y: int)
    +Draw()
}

class ImageEditor {
    -all: CompoundGraphic
    +Load()
    +GroupSelected(components: []Graphic)
}

Graphic <|.. Dot
Graphic <|.. Circle
Graphic <|.. CompoundGraphic

CompoundGraphic "1" --> "*" Graphic
ImageEditor "1" --> "1" CompoundGraphic
```