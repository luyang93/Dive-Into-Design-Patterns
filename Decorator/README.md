```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface Coffee {
  +Cost() : float64
}

class SimpleCoffee {
  +Cost() : float64
}

class MilkDecorator {
  -coffee : Coffee
  +Cost() : float64
}

class SugarDecorator {
  -coffee : Coffee
  +Cost() : float64
}

Coffee <|.. SimpleCoffee
Coffee <|.. MilkDecorator
Coffee <|.. SugarDecorator

MilkDecorator o-- Coffee: has
SugarDecorator o-- Coffee: has
```