```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface Logger {
    +SetNext(Logger)
    +LogMessage(int, string)
}

class AbstractLogger {
    -Level : int
    -NextLogger : Logger
    -WriteFunc : func(string)
    +SetNext(Logger)
    +LogMessage(int, string)
}

class ErrorLogger {
    -AbstractLogger
    +Write(string)
}

class DebugLogger {
    -AbstractLogger
    +Write(string)
}

class InfoLogger {
    -AbstractLogger
    +Write(string)
}

Logger <|.. AbstractLogger
AbstractLogger <|-- ErrorLogger
AbstractLogger <|-- DebugLogger
AbstractLogger <|-- InfoLogger
```