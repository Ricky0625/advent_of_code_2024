package set

// Set implementation using map in go
type Set[T comparable] struct {
	elements map[T]struct{}
}

// returns a new Set
func New[T comparable]() *Set[T] {
	// create a pointer to a new instance of a generic struct Set[T]
	return &Set[T]{elements: make(map[T]struct{})}
}

// add new element
func (s *Set[T]) Add(value T) {
	s.elements[value] = struct{}{}
}

// remove an element
func (s *Set[T]) Remove(value T) {
	delete(s.elements, value)
}

// check if value exists
func (s *Set[T]) Contains(value T) bool {
	_, exists := s.elements[value]
	return exists
}

// return keys as a slice of T
func (s *Set[T]) List() []T {
	var list []T
	for key := range s.elements {
		list = append(list, key)
	}
	return list
}
