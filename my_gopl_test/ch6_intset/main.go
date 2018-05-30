package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// Remove remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word >= len(s.words) {
		return
	}
	s.words[word] &= ^(1 << bit)
}

// Clear remove all elements from the set
func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

// Len return the number of elements
func (s *IntSet) Len() int {
	var len int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				len++
			}
		}
	}
	return len
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Copy return a copy of the set
func (s *IntSet) Copy() *IntSet {
	n := new(IntSet)
	for _, word := range s.words {
		n.words = append(n.words, word)
	}
	return n
}

// String returns the set as a string of te form "{1 2 3}"
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"
	fmt.Println(x.Len())    // 3

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"
	fmt.Println(y.Len())    // 2

	x.UnionWith(&y)
	fmt.Println(x.String())           // "{1 9 42 144}"
	fmt.Println(x.Len())              // 4
	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)

	z := x.Copy()
	fmt.Println(z) // "{1 9 42 144}"

	x.Remove(2)
	fmt.Println(&x) // "{1 9 42 144}"
	x.Remove(42)
	fmt.Println(&x) // "{1 9 144}"
	x.Remove(145)
	fmt.Println(&x) // "{1 9 144}"
	x.Remove(144)
	fmt.Println(&x) // "{1 9}"

	fmt.Println(x.Len()) // 2
	fmt.Println(&x)      // "{1 9}"
	x.Clear()
	fmt.Println(x.Len()) // 0
	fmt.Println(&x)      // "{}"

	fmt.Println(z) // "{1 9 42 144}"
}
