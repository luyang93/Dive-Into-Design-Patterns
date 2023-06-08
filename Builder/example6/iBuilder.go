package main

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumberFloor()
	getHouse() house
}

func getBuilder(builderType string) iBuilder {
	switch builderType {
	case "normal":
		return newNormalBuilder()
	case "igloo":
		return newIglooBuilder()
	}
	return nil
}
