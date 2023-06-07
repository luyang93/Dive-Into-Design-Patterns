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
		return &normalBuilder{}
	case "igloo":
		return &iglooBuilder{}
	}
	return nil
}
