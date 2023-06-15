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
```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface Department {
    +execute(patient: Patient): void
    +setNext(next: Department): void
}

class Reception {
    -next: Department
    +execute(patient: Patient): void
    +setNext(next: Department): void
}

class Doctor {
    -next: Department
    +execute(patient: Patient): void
    +setNext(next: Department): void
}

class Medical {
    -next: Department
    +execute(patient: Patient): void
    +setNext(next: Department): void
}

class Cashier {
    -next: Department
    +execute(patient: Patient): void
    +setNext(next: Department): void
}

class Patient {
    -name: string
    -registrationDone: bool
    -doctorCheckUpDone: bool
    -medicineDone: bool
    -paymentDone: bool
}

Department <|.. Reception
Department <|.. Doctor
Department <|.. Medical
Department <|.. Cashier

Reception "1" --> "1" Department
Doctor "1" --> "1" Department
Medical "1" --> "1" Department
Cashier "1" --> "1" Department

Patient "1" --> "*" Department
```
```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface DispenseChain {
  +SetNextChain(nextChain: DispenseChain)
  +Dispense(currency: Currency)
}

class Dollar50Dispenser {
  +Chain: DispenseChain
  +SetNextChain(nextChain: DispenseChain)
  +Dispense(currency: Currency)
}

class Dollar20Dispenser {
  +Chain: DispenseChain
  +SetNextChain(nextChain: DispenseChain)
  +Dispense(currency: Currency)
}

class Dollar10Dispenser {
  +Chain: DispenseChain
  +SetNextChain(nextChain: DispenseChain)
  +Dispense(currency: Currency)
}

class Currency {
  -Amount: int
}

class ATMDispenser {
  -C1: DispenseChain
  +DispenseMoney(currency: Currency)
}

DispenseChain <|.. Dollar50Dispenser
DispenseChain <|.. Dollar20Dispenser
DispenseChain <|.. Dollar10Dispenser

Dollar50Dispenser o-- Dollar20Dispenser: "chain"
Dollar20Dispenser o-- Dollar10Dispenser: "chain"

ATMDispenser o-- Dollar50Dispenser: "C1"
ATMDispenser o-- Currency
```