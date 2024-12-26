package main

type Set map[int]bool

func set() Set {
	return Set{}
}

func (s Set) add(i int) {
	s[i] = true
}

func (s Set) has(i int) bool {
	_, ok := s[i]
	return ok
}
