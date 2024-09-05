package main

import (
	"fmt"
	"iter"
	"iterators/set"
	"iterators/tree"
)

func PrintAllElements[E any](s iter.Seq[E]) {
	for v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func EqSeq[E comparable](s1, s2 iter.Seq[E]) bool {
	next1, stop1 := iter.Pull(s1)
	defer stop1()
	next2, stop2 := iter.Pull(s2)
	defer stop2()

	for {
		v1, ok1 := next1()
		v2, ok2 := next2()
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if !ok1 {
			return true
		}
	}
}

func Filter[E any](pred func(E) bool, s iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range s {
			if pred(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func main() {
	s := set.New[int]()
	s.Add(3)
	s.Add(7)
	s.Add(11)
	s.Add(-1)

	set.PrintAllElementsPush(s)
	set.PrintAllElementsPull(s)

	PrintAllElements(s.All())

	t := tree.New[int](5)
	t.Insert(7)
	t.Insert(11)
	t.Insert(-1)

	t2 := tree.New(5)

	PrintAllElements(t.All())

	fmt.Println(EqSeq(t.All(), t.All()))
	fmt.Println(EqSeq(t.All(), t2.All()))

	PrintAllElements(Filter(func(e int) bool {
		return e > 0
	}, t.All()))
}
