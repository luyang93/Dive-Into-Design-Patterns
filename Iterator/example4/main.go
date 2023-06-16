package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s *Student) GetName() string {
	return s.Name
}

func (s *Student) GetAge() int {
	return s.Age
}

func NewStudent(Name string, Age int) *Student {
	return &Student{
		Name: Name,
		Age:  Age,
	}
}

type Iterator interface {
	HasNext() bool
	Next() *Student
}

type ClassroomIterator struct {
	Index     int
	Classroom *Classroom
}

func (c *ClassroomIterator) HasNext() bool {
	return c.Index < c.Classroom.GetSize()
}

func (c *ClassroomIterator) Next() *Student {
	student := c.Classroom.GetStudentAt(c.Index)
	c.Index++
	return student
}

type Classroom struct {
	Students []*Student
}

func (c *Classroom) AddStudent(student *Student) {
	c.Students = append(c.Students, student)
}

func (c *Classroom) GetStudentAt(Index int) *Student {
	return c.Students[Index]
}

func (c *Classroom) GetSize() int {
	return len(c.Students)
}

func (c *Classroom) CreateIterator() Iterator {
	return &ClassroomIterator{
		Index:     0,
		Classroom: c,
	}
}

func NewClassroom() *Classroom {
	return &Classroom{}
}

func main() {
	classroom := NewClassroom()
	classroom.AddStudent(&Student{"Alice", 20})
	classroom.AddStudent(&Student{"Bob", 21})
	classroom.AddStudent(&Student{"Charlie", 22})

	iterator := classroom.CreateIterator()
	for iterator.HasNext() {
		student := iterator.Next()
		fmt.Println(student.Name, ":", student.Age)
	}
}
