```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

class RoundHole {
    -radius: float64
    +GetRadius(): float64
    +Fits(peg: Peg): bool
}

interface Peg {
    +GetRadius(): float64
}

class RoundPeg {
    -radius: float64
    +GetRadius(): float64
}

class SquarePeg {
    -width: float64
    +GetWidth(): float64
}

class SquarePegAdapter {
    +GetRadius(): float64
}

RoundHole --> Peg : uses >
RoundPeg ..|> Peg : implements >
SquarePegAdapter ..|> Peg : implements >
SquarePegAdapter --> SquarePeg : adapts >
```
```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface "Computer" {
    +InsertIntoLightningPort()
}

class Client {
    +InsertLightningConnectorIntoComputer(computer : Computer)
}

class Mac {
    +InsertIntoLightningPort()
}

class Windows {
    +InsertIntoUSBPort()
}

class WindowsAdapter {
    +windowComputer: Windows
    +InsertIntoLightningPort()
}

"Client" --> "Computer" : uses >
"Mac" ..|> "Computer"
"WindowsAdapter" ..|> "Computer"
"WindowsAdapter" --> "Windows" : adapts >
```