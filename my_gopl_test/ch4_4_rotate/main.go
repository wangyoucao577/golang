package main

import (
	"fmt"
)

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Rotate s left by num positions
func rotate(s []int, num int) {
	reverse(s[:num])
	reverse(s[num:])
	reverse(s)
}

func rotate2(s []int, num int) {
	out := make([]int, len(s))
	copy(out, s)

	for i, v := range s {
		if i < num {
			out = out[1:]
			out = append(out, v)
		}
	}

	copy(s, out)
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(a)

	reverse(a[:])
	fmt.Println(a)

	rotate(a[:], 2)
	fmt.Println(a)

	rotate2(a[:], 2)
	fmt.Println(a)

}
