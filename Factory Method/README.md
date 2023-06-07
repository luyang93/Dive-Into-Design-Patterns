```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface iGun {
    + setName(string)
    + setPower(int)
    + getName() string
    + getPower() int
}

class gun {
    + name: string
    + power: int
    + setName(string)
    + setPower(int)
    + getName() string
    + getPower() int
}

class ak47 {
    + gun
}

class maverick {
    + gun
}

class GunFactory {
    + getGun(string) iGun
}

class Client

ak47 <|.. iGun: implementation
maverick <|.. iGun: implementation
ak47 o--> gun: aggregation
maverick o--> gun: aggregation
GunFactory --> iGun: creates
Client --> GunFactory
```