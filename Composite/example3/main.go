package main

import "fmt"

type Component interface {
	PrintName()
}

type Employee struct {
	name string
}

func (e Employee) PrintName() {
	fmt.Println("\t\t\tEmployee: ", e.name)
}

type Department struct {
	name       string
	components []Component
}

func (d Department) PrintName() {
	fmt.Println("\t\tDepartment: ", d.name)
	for _, component := range d.components {
		component.PrintName()
	}
}

type Team struct {
	name       string
	components []Component
}

func (t Team) PrintName() {
	fmt.Println("\tTeam: ", t.name)
	for _, component := range t.components {
		component.PrintName()
	}
}

type Manager struct {
	name       string
	components []Component
}

func (m Manager) PrintName() {
	fmt.Println("Manager: ", m.name)
	for _, component := range m.components {
		component.PrintName()
	}
}

func main() {
	employees := []Component{
		Employee{name: "John"},
		Employee{name: "Jane"},
		Employee{name: "Bob"},
	}

	department := Department{
		name:       "Engineering",
		components: employees,
	}
	team := Team{
		name:       "Product",
		components: []Component{department},
	}
	manager := Manager{
		name:       "Alice",
		components: []Component{team, team},
	}

	manager.PrintName()
}
