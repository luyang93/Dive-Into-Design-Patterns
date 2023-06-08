```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface iGun {
    + setName(string)
    + setPower(int)
    + getName() string
    + getPower() int
}

class Gun {
    + name: string
    + power: int
    + setName(string)
    + setPower(int)
    + getName() string
    + getPower() int
}

class ak47 {
    + Gun
}

class maverick {
    + Gun
}

class GunFactory {
    + getGun(string) iGun
}

class Client

Gun ..|> iGun

ak47 --|> Gun
maverick --|> Gun

GunFactory ..> ak47
GunFactory ..> maverick

Client --> GunFactory
```
```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface Transport {
    + Deliver() string
}

class Ship {
    + Deliver() string
}

class Trunk {
    + Deliver() string
}

interface Logistics {
    + createTransport() Transport
    + planDelivery() string
}

class RoadLogistics {
    + createTransport() Transport
    + planDelivery() string
}

class SeaLogistics {
    + createTransport() Transport
    + planDelivery() string
}

class LogisticsFactory {
    createLogistics() Logistics
}

class client

Ship ..|> Transport
Trunk ..|> Transport

RoadLogistics ..|> Logistics
SeaLogistics ..|> Logistics

RoadLogistics ..> Trunk
SeaLogistics ..> Ship

LogisticsFactory ..> RoadLogistics
LogisticsFactory ..> SeaLogistics

client --> LogisticsFactory
```