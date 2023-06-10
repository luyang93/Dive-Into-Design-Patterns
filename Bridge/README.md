```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

class RemoteControl {
  device: Device
  TogglePower()
  volumeDown()
  volumeUp()
  channelDown()
  channelUp()
}

class AdvancedRemoteControl {
  Mute()
}

interface Device {
  IsEnabled(): bool
  Enable()
  Disable()
  GetVolume(): int
  SetVolume(percent: int)
  GetChannel(): int
  SetChannel(channel: int)
}

class TV {
  isEnabled: bool
  volume: int
  channel: int
  IsEnabled(): bool
  Enable()
  Disable()
  GetVolume(): int
  SetVolume(volume: int)
  GetChannel(): int
  SetChannel(channel: int)
}

class Radio {
  isEnabled: bool
  volume: int
  channel: int
  IsEnabled(): bool
  Enable()
  Disable()
  GetVolume(): int
  SetVolume(volume: int)
  GetChannel(): int
  SetChannel(channel: int)
}

Device <|.. TV
Device <|.. Radio
RemoteControl o-- Device
AdvancedRemoteControl --|> RemoteControl
```
```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface Painter {
    + Paint(s: Shape)
}

class RedPainter {
    + Paint(s: Shape)
}

class BluePainter {
    + Paint(s: Shape)
}

class Shape {
    - shape: string
    - color: string
    - painter: Painter

    + Draw(): void
    + SetPainter(painter: Painter)
}

class Circle {
    - shape: string
    - color: string
    - painter: Painter
}

class Square {
    - shape: string
    - color: string
    - painter: Painter
}

Painter <|.. RedPainter
Painter <|.. BluePainter
Shape "1" o--> "1" Painter
Circle --|> Shape
Square --|> Shape
```
```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface Payment {
  +ProcessPayment()
}

class CreditCardPayment {
  +ProcessPayment()
}

class WeChatPayment {
  +ProcessPayment()
}

class CashPayment {
  +ProcessPayment()
}

interface Product {
  +Purchase()
}

class Book {
  +Purchase()
}

class Electronics {
  +Purchase()
}

Payment <|.. CreditCardPayment
Payment <|.. WeChatPayment
Payment <|.. CashPayment

Product <|.. Book
Product <|.. Electronics

Book ..> Payment : uses
Electronics ..> Payment : uses
```