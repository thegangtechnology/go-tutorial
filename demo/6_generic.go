package demo

import "fmt"

type Stack[T any] struct {
	Items []T
}

func (s *Stack[T]) Push(item T) {
	s.Items = append(s.Items, item)
}

func (s *Stack[T]) Pop() T {
	if len(s.Items) == 0 {
		var zero T
		return zero
	}
	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return item
}

func StackDemo() {
	var s Stack[int]
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	var s2 Stack[string]
	s2.Push("a")
	s2.Push("b")
	s2.Push("c")
	fmt.Println(s2.Pop())
	fmt.Println(s2.Pop())
	fmt.Println(s2.Pop())
}
