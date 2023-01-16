package main

import "fmt"

type Stack[T comparable] struct {
	elements []T
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s Stack[T]) Size() int {
	return len(s.elements)
}

func (s Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zeroVal T
		return zeroVal, false
	}
	popped := s.elements[s.Size()-1]
	s.elements = s.elements[:len(s.elements)-1]
	return popped, true
}

func (s Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zeroVal T
		return zeroVal, false
	}
	return s.elements[len(s.elements)-1], true
}

func main() {
	var s Stack[int]
	s.Push(143)
	s.Push(153)
	fmt.Println(s)
	fmt.Println("Size of stack is ", s.Size())
	i, ok := s.Peek()
	if ok {
		fmt.Println("Top element of stack is ", i)
	}
	s.Pop()
	fmt.Println("Size of stack is ", s.Size())
	s.Pop()
	fmt.Println("Is stack empty ", s.IsEmpty())

}
