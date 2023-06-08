```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface inode {
    + clone() inode
    + print(string)
} 

class file {
    + name string
    + clone() inode
    + print(string)
}

class folder {
    + name string
    + children []inode
    + clone() inode
    + print(string)
}

class client {
    + inode.clone() inode
}

file ..|> inode
folder ..|> inode

client --> inode
```
```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface Cloneable {
    + Clone() Cloneable
}

class Shape {
    - x int
    - y int
    - color string
    + Shape()
    + Clone()
}

class Rectangle {
    - width int
    - height int
    + Rectangle()
    + Clone()
}

class Circle {
    - radius int
    + Circle()
    + Clone()
}

Shape ..|> Cloneable

Rectangle --|> Shape
Circle --|> Shape

client o--> Shape
```
```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface Prototype {
    +Clone(): Prototype
    +Use(): string
}

class Registry {
    +prototypes: map[string]Prototype
    +NewRegistry(): *Registry
    +Register(name: string, prototype: Prototype)
    +Create(name: string): (Prototype, error)
}

class ConcretePrototype {
    +data: string
    +NewConcretePrototype(data: string): Prototype
    +Clone(): Prototype
    +Use(): string
}

Prototype <|.. ConcretePrototype
Registry o--> Prototype
client --> Registry
```