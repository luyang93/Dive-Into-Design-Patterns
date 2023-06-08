```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface IShirt {
    + setLogo(string)
    + setSize(int)
    + getLogo() string
    + getSize() int
}

class Shirt{
    + logo string
    + size int
    + setLogo(string)
    + setSize(int)
    + getLogo() string
    + getSize() int
}

interface IShoe {
    + setLogo(string)
    + setSize(int)
    + getLogo() string
    + getSize() int
}

class Shoe{
    + logo string
    + size int
    + setLogo(string)
    + setSize(int)
    + getLogo() string
    + getSize() int
}

interface ISportFactory {
    + makeShoe() IShoe
    + makeShirt() IShirt
}

class Adidas {
    + makeShoe() IShoe
    + makeShirt() IShirt
}

class Nike {
    + makeShoe() IShoe
    + makeShirt() IShirt
}

class AdidasShoe {
    + Shoe
}

class NikeShoe {
    + Shoe
}

class AdidasShirt {
    + Shirt
}

class NikeShirt {
    + Shirt
}

class Client {
    + GetSportsFactory(int) ISportFactory
}


Shoe ..|> IShoe
Shirt ..|> IShirt

Adidas ..|> ISportFactory
Nike ..|> ISportFactory

NikeShoe --|> Shoe
AdidasShoe --|> Shoe

NikeShirt --|> Shirt
AdidasShirt --|> Shirt

Adidas ..> AdidasShirt
Adidas ..> AdidasShoe

Nike ..> NikeShirt
Nike ..> NikeShoe

Client --> ISportFactory
```