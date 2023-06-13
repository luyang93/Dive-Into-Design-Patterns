package main

import "fmt"

type SoldierFlyweight interface {
	Display(x, y int)
}

type SoldierType struct {
	Type, Weapon, Armor string
}

func NewSoldierType(Type, Weapon, Armor string) *SoldierType {
	return &SoldierType{Type, Weapon, Armor}
}

func (s *SoldierType) Display(x int, y int) {
	fmt.Printf("Displaying %s at (%d, %d) with weapon %s and armor %s\n", s.Type, x, y, s.Weapon, s.Armor)
}

type SoldierFlyweightFactory struct {
	soldierTypeMap map[string]*SoldierType
}

func (sf *SoldierFlyweightFactory) GetSoldierType(Type string, Weapon string, Armor string) *SoldierType {
	soldierType, ok := sf.soldierTypeMap[Type]
	if !ok {
		soldierType = NewSoldierType(Type, Weapon, Armor)
		sf.soldierTypeMap[Type] = soldierType
	}
	return soldierType
}

type Soldier struct {
	soldierFlyweight SoldierFlyweight
	x                int
	y                int
}

func (s *Soldier) Display() {
	s.soldierFlyweight.Display(s.x, s.y)
}

func main() {
	soldierTypeFactory := SoldierFlyweightFactory{make(map[string]*SoldierType)}

	archerType := soldierTypeFactory.GetSoldierType("Archer", "Bow", "Leather")
	knightType := soldierTypeFactory.GetSoldierType("Knight", "Sword", "Plate")

	archer1 := Soldier{archerType, 100, 50}
	archer2 := Soldier{archerType, 120, 60}
	knight1 := Soldier{knightType, 200, 100}
	knight2 := Soldier{knightType, 250, 120}

	archer1.Display()
	archer2.Display()
	knight1.Display()
	knight2.Display()
}
