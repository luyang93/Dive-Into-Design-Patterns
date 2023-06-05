package main

import "fmt"

type Object struct {
	// Define the properties of the object
	Param1 string
	Param2 int
	Param3 bool
}

type ObjectBuilder struct {
	obj Object
}

func NewObjectBuilder() *ObjectBuilder {
	return &ObjectBuilder{}
}

// Set the first optional parameter
func (b *ObjectBuilder) WithParam1(param1 string) *ObjectBuilder {
	b.obj.Param1 = param1
	return b
}

// Set the second optional parameter
func (b *ObjectBuilder) WithParam2(param2 int) *ObjectBuilder {
	b.obj.Param2 = param2
	return b
}

// Set the third optional parameter
func (b *ObjectBuilder) WithParam3(param3 bool) *ObjectBuilder {
	b.obj.Param3 = param3
	return b
}

// Build the object
func (b *ObjectBuilder) Build() *Object {
	return &b.obj
}

func main() {
	// Use the builder pattern to construct the object
	objBuilder := NewObjectBuilder()
	obj := objBuilder.WithParam1("value1").WithParam2(44).WithParam3(true).Build()
	fmt.Println(obj)
}
