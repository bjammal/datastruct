//Package stack provides a simple implementation of a stack
package stack

//Stack defines a stack data structure
type Stack struct {
	elements []interface{}
}

//Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return len(s.elements)
}

//Pop removes the element on the top of the stack and returns it
func (s *Stack) Pop() interface{} {
	e := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return e
}

//Push adds an element to the top of the stack
func (s *Stack) Push(e interface{}) {
	s.elements = append(s.elements, e)
}

//Peek returns the element on the top of the stack without removing it
func (s *Stack) Peek() interface{} {
	return s.elements[len(s.elements)-1]
}

//IsEmpty return true if the stack is empty, false otherwise
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}
