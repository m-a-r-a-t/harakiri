/*Реализовать пересечение двух неупорядоченных множеств.*/
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

func NewSet[T comparable]() Set[T] {
	return Set[T]{
		data: make(map[T]struct{}),
	}
}

func main() {
	setOne := NewSet[int64]()
	setTwo := NewSet[int64]()

	setOne.Add(1, 35, 62, 14, 99)
	setTwo.Add(5, 35, 99, 111, 1)

	fmt.Println(setOne.Intersection(setTwo))

	

}
