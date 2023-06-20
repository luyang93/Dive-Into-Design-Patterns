```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface Strategy {
  +execute(a, b int) int
}

class Addition << (S, #ADD1B2) Strategy >> {
  +execute(a, b int) int
}

class Subtraction << (S, #ADD1B2) Strategy >> {
  +execute(a, b int) int
}

class Multiplication << (S, #ADD1B2) Strategy >> {
  +execute(a, b int) int
}

class Context {
  -strategy Strategy
  +SetStrategy(s Strategy)
  +ExecuteStrategy(a, b int) int
}

Strategy <|.. Addition : implements
Strategy <|.. Subtraction : implements
Strategy <|.. Multiplication : implements

Context o-- Strategy : uses >
```