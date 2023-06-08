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