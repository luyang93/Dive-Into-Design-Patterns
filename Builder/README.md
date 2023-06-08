```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface iBuilder {
    + setWindowType()
	+ setDoorType()
	+ setNumberFloor()
	+ getHouse() house
}

class normalBuilder {
	+ windowType string
	+ doorType string
	+ floor int
    + setWindowType()
	+ setDoorType()
	+ setNumberFloor()
	+ getHouse() house
}

class iglooBuilder {
	+ windowType string
	+ doorType string
	+ floor int
    + setWindowType()
	+ setDoorType()
	+ setNumberFloor()
	+ getHouse() house
}


class house {
    + windowType string
	+ doorType string
	+ floor int
}

class director {
    + builder iBuilder
    + setBuilder(iBuilder)
    + buildHouse() house
}

class client {
    + getBuilder(string) iBuilder
    + newDirector(string) director
}

iglooBuilder ..|> iBuilder
normalBuilder ..|> iBuilder

iglooBuilder --> house
normalBuilder --> house

director --> iBuilder

client --> director
client ..> iBuilder

```