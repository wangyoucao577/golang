package main

import (
	"fmt"
)

func removeDuplicateString(s []string) []string {
	out := s
	var count int

	for _, v := range s {
		exist := false
		for _, v1 := range out[:count] {
			if v1 == v {
				exist = true
				break
			}
		}

		if !exist {
			out[count] = v
			count++
		}
	}

	return out[:count]
}

func main() {
	s := [...]string{"test", "test", "test1", "test2", "test3", "test2"}

	s1 := s[:]
	fmt.Printf("len: %d ,cap: %d, slice: %v\n", len(s1), cap(s1), s1)

	s1 = removeDuplicateString(s1)
	fmt.Printf("len: %d ,cap: %d, slice: %v\n", len(s1), cap(s1), s1)
}
