//Package queue provides a simple implementation of a Queue
package queue

//Queue defines a queue data structure
type Queue struct {
	elements []interface{}
}

//Size returns the number of elements in the queue
func (s *Queue) Size() int {
	return len(s.elements)
}

//Remove removes the element on the top of the queue and returns it
func (s *Queue) Remove() interface{} {
	e := s.elements[0]
	s.elements = s.elements[1:]
	return e
}

//Add adds an element to the top of the queue
func (s *Queue) Add(e interface{}) {
	s.elements = append(s.elements, e)
}

//Peek returns the element on the top of the queue without removing it
func (s *Queue) Peek() interface{} {
	return s.elements[0]
}

//IsEmpty return true if the queue is empty, false otherwise
func (s *Queue) IsEmpty() bool {
	return len(s.elements) == 0
}
