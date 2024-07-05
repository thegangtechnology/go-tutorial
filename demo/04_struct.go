package demo

import "fmt"

type Student struct {
	Name string
	GPA  float64
}

type ClassRoom struct {
	RoomName string
	Students []Student
}

// There is no constructor in golang
// but you can do this
func EmptyClassRoom(name string) ClassRoom {
	return ClassRoom{
		RoomName: name,
		Students: make([]Student, 0),
	}
}

func (c *ClassRoom) AddStudent(s Student) {
	c.Students = append(c.Students, s)
}

func (c *ClassRoom) NumStudents() int {
	return len(c.Students)
}

func InGoAssignmentIsCopy() {
	a := Student{Name: "A", GPA: 1.23}

	//class1 := ClassRoom{RoomName: "Class1"}
	class1 := EmptyClassRoom("Class1")
	class2 := class1
	class1.AddStudent(a)

	fmt.Printf("class1: %v\n", class1)
	fmt.Printf("class2: %v\n", class2)
}

func UsePointerIfYouMeantIt() {
	a := Student{Name: "A", GPA: 1.23}
	//var class1 *ClassRoom
	class1 := &ClassRoom{RoomName: "Class1"}
	class2 := class1
	class1.AddStudent(a)

	fmt.Printf("class1: %v\n", class1)
	fmt.Printf("class2: %v\n", class2)
}
