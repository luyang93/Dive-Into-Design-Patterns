```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface Expression {
  +Interpret()
}

class PrintExpression {
  -Message: string
  +Interpret(): void
}

class RepeatExpression {
  -RepeatCount: int
  -Expression: Expression
  +Interpret(): void
}

class ReverseExpression {
  -Message: string
  +Interpret(): void
}

PrintExpression -up-|> Expression
RepeatExpression -up-|> Expression
ReverseExpression -up-|> Expression
```