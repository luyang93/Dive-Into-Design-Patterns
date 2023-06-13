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
```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface SoldierFlyweight {
    +Display(x: int, y: int)
}

class SoldierType {
    -Type: string
    -weapon: string
    -armor: string
    +Display(x: int, y: int)
}

class SoldierFlyweightFactory {
    -soldierTypeMap: map[string]*SoldierType
    +GetSoldierType(Type: string, weapon: string, armor: string): *SoldierType
}

class Soldier {
    -soldierFlyweight: SoldierFlyweight
    -x: int
    -y: int
    +Display()
}

SoldierType .up.|> SoldierFlyweight
Soldier .right.> SoldierFlyweight: Uses >
SoldierFlyweightFactory .down.> SoldierType: Creates >
```