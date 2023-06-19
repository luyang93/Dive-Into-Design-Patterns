```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

class TextEditor {
  -content: string
  +write(text: string): void
  +save(): Memento
  +restore(m: Memento): void
}

class Memento {
  -content: string
  +getContent(): string
}

TextEditor "1" -- "1" Memento : creates >
```
```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

class Calculator {
  -result: int
  +add(value: int)
  +subtract(value: int)
  +save(): CalculatorMemento
  +restore(m: CalculatorMemento)
  +getResult(): int
}

class CalculatorMemento {
  -result: int
}

class CalculatorHistory {
  -history: []CalculatorMemento
  +save(c: Calculator)
  +undo(c: Calculator)
}

Calculator -right-> CalculatorMemento: <<creates>>
Calculator o-down-> CalculatorMemento: <<uses>>
CalculatorHistory o-left-> Calculator: <<uses>>
```