package demo

import "fmt"

type HasStudents interface {
	AddStudent(s Student)
	NumStudents() int
}

type SchoolBus struct {
	Color    string
	Students []Student
}

func (s *SchoolBus) AddStudent(st Student) {
	s.Students = append(s.Students, st)
}

func (s *SchoolBus) NumStudents() int {
	return len(s.Students)
}

func Add2Students(s HasStudents) {
	s.AddStudent(Student{Name: "A", GPA: 1.23})
	s.AddStudent(Student{
		Name: "B",
		GPA:  3.45,
	})
}

func InterfaceDemo() {
	s := &SchoolBus{Color: "red"}
	Add2Students(s)
	fmt.Println("School Bus", s.NumStudents())

	// note how we never specify classroom implements HasStudents
	// go is compile-time ducktyping
	c := &ClassRoom{RoomName: "Class1"}
	fmt.Println("Class Room", c.NumStudents())
}
