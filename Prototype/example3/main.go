package main

import "fmt"

type Prototype interface {
	Clone() Prototype
	Use() string
}

type Registry struct {
	prototypes map[string]Prototype
}

func NewRegistry() *Registry {
	return &Registry{
		make(map[string]Prototype),
	}
}

func (r *Registry) Register(name string, prototype Prototype) {
	r.prototypes[name] = prototype
}

func (r *Registry) Create(name string) (Prototype, error) {
	if prototype, found := r.prototypes[name]; found {
		return prototype.Clone(), nil
	} else {
		return nil, fmt.Errorf("Prototype not found: %s", name)
	}
}

type ConcretePrototype struct {
	data string
}

func NewConcretePrototype(data string) Prototype {
	return &ConcretePrototype{
		data,
	}
}

func (p *ConcretePrototype) Clone() Prototype {
	return NewConcretePrototype(p.data)
}

func (p *ConcretePrototype) Use() string {
	return p.data
}

func main() {
	registry := NewRegistry()

	registry.Register("prototype1", NewConcretePrototype("Data for prototype1"))
	registry.Register("prototype2", NewConcretePrototype("Data for prototype2"))

	clone1, _ := registry.Create("prototype1")
	clone2, _ := registry.Create("prototype2")
	clone3, _ := registry.Create("prototype2")

	fmt.Println(clone1.Use()) // "Data for prototype1"
	fmt.Println(clone2.Use()) // "Data for prototype2"

	fmt.Println(clone3 == clone2)
}
