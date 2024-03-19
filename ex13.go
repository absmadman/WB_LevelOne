package main

import "fmt"

type Set struct {
	set map[string]int
}

func NewSet() *Set {
	return &Set{
		set: make(map[string]int),
	}
}

func (s *Set) Insert(elem string) {
	if _, ok := s.set[elem]; ok {
		s.set[elem]++
	} else {
		s.set[elem] = 1
	}
}

func (s *Set) Contains(elem string) bool {
	if _, ok := s.set[elem]; ok {
		return true
	}
	return false
}

func (s *Set) Delete(elem string) {
	if val, ok := s.set[elem]; ok {
		if val > 1 {
			s.set[elem]--
		} else {
			delete(s.set, elem)
		}
	}
}

func (s *Set) Erase(elem string) {
	if s.Contains(elem) {
		delete(s.set, elem)
	}
}

func (s *Set) Count(elem string) int {
	if val, ok := s.set[elem]; ok {
		return val
	}
	return 0
}

func main() {
	set := NewSet()
	n := []string{"cat", "cat", "dog", "cat", "tree"}
	for i := range n {
		set.Insert(n[i])
	}
	fmt.Println(set.Count("cat"))
	set.Delete("cat")

	fmt.Println(set.Count("cat"))
	set.Erase("cat")

	fmt.Println(set.Count("cat"))
}
