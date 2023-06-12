```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface Coffee {
    +Cost(): float64
}

class SimpleCoffee {
    +Cost(): float64
}

class CoffeeDecorator {
    +WithMilk(coffee: Coffee): Coffee
    +WithSugar(coffee: Coffee): Coffee
}

class MilkDecorator {
    -coffee: Coffee
    +Cost(): float64
}

class SugarDecorator {
    -coffee: Coffee
    +Cost(): float64
}

Coffee <|.. SimpleCoffee
Coffee <|.. MilkDecorator
Coffee <|.. SugarDecorator
CoffeeDecorator o-- Coffee
MilkDecorator --> CoffeeDecorator
SugarDecorator --> CoffeeDecorator
```