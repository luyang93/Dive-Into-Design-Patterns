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