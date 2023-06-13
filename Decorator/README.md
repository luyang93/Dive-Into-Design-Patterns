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
```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface DataSource {
  +WriteData(data: string)
  +ReadData(): string
}

class FileDataSource {
  -filename: string
  -data: string
  +WriteData(data: string)
  +ReadData(): string
}

class DataSourceDecorator {
  -wrappee: DataSource
  +WriteData(data: string)
  +ReadData(): string
}

class EncryptionDecorator {
  +WriteData(data: string)
  +ReadData(): string
}

class CompressionDecorator {
  +WriteData(data: string)
  +ReadData(): string
}

DataSource <|.. FileDataSource
DataSource <|.. DataSourceDecorator
DataSourceDecorator <|-- EncryptionDecorator
DataSourceDecorator <|-- CompressionDecorator
DataSourceDecorator o-- DataSource
```
```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface Notifier {
    + Send(message: String)
}

class EmailNotifier {
    + Send(message: String)
}

class NotifierDecorator {
    - notifier: Notifier
    + Send(message: String)
}

class WeChatDecorator {
    - notifier: NotifierDecorator
    + Send(message: String)
}

class MobileDecorator {
    - notifier: NotifierDecorator
    + Send(message: String)
}

Notifier <|.. EmailNotifier
Notifier <|.. NotifierDecorator
NotifierDecorator <|-- WeChatDecorator
NotifierDecorator <|-- MobileDecorator
```