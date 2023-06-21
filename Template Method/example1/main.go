package main

import "fmt"

type Position struct {
	CoordinateX int
	CoordinateY int
}

type Structure struct {
	Resource int
}

type GameAI interface {
	Turn()
	CollectResources()
	BuildStructures()
	BuildUnits()
	Attack()
	SendScouts(Position)
	SendWarriors(Position)
}

type BaseGameAI struct {
	BuiltStructures []Structure
}

func (b *BaseGameAI) Turn() {
	b.CollectResources()
	b.BuildStructures()
	b.BuildUnits()
	b.Attack()
}

func (b *BaseGameAI) CollectResources() {
	for _, s := range b.BuiltStructures {
		s.Resource++
	}
}

// Placeholder methods for BaseGameAI
func (b *BaseGameAI) BuildStructures()               {}
func (b *BaseGameAI) BuildUnits()                    {}
func (b *BaseGameAI) Attack()                        {}
func (b *BaseGameAI) SendScouts(position Position)   {}
func (b *BaseGameAI) SendWarriors(position Position) {}

type OrcsAI struct {
	BaseGameAI
	Resources int
	Scouts    []int
	Warriors  []int
}

func (o *OrcsAI) BuildStructures() {
	if o.Resources > 0 {
		// Build farms, then barracks, then stronghold.
	}
}

func (o *OrcsAI) BuildUnits() {
	if o.Resources > 5 {
		if len(o.Scouts) == 0 {
			// Build peon, add it to scouts group.
		} else {
			// Build grunt, add it to warriors group.
		}
	}
}

func (o *OrcsAI) SendScouts(position Position) {
	if len(o.Scouts) > 0 {
		// Send scouts to position.
	}
}

func (o *OrcsAI) SendWarriors(position Position) {
	if len(o.Warriors) > 5 {
		// Send warriors to position.
	}
}

type MonstersAI struct {
	BaseGameAI
}

func (m *MonstersAI) CollectResources() {
	// Monsters don't collect resources.
}

func (m *MonstersAI) BuildStructures() {
	// Monsters don't build structures.
}

func (m *MonstersAI) BuildUnits() {
	// Monsters don't build units.
}

func main() {
	orcAI := &OrcsAI{}
	orcAI.Turn()

	monstersAI := &MonstersAI{}
	monstersAI.Turn()

	fmt.Println("Game AI implemented successfully!")
}
