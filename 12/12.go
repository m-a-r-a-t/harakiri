package main

import "fmt"

// структура множество
type Set[T comparable] struct {
	data map[T]struct{}
}

// метод добавления значений в множество
func (s *Set[T]) Add(values ...T) {
	for _, val := range values {
		if _, ok := s.data[val]; !ok {
			s.data[val] = struct{}{}
		}
	}
}

// метод поиска пересечений
func (s *Set[T]) Intersection(otherSet Set[T]) Set[T] {
	var newSet = Set[T]{
		data: make(map[T]struct{}),
	}

	for value, _ := range s.data {
		if _, ok := otherSet.data[value]; ok {
			newSet.Add(value)
		}
	}

	return newSet
}

func (s *Set[T]) GetValues() map[T]struct{} {

	return s.data
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{
		data: make(map[T]struct{}),
	}
}

func main() {
	setOneStr := NewSet[string]()

	setOneStr.Add("cat", "cat", "dog", "cat", "tree")

	fmt.Println(setOneStr.GetValues())
}
