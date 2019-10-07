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
	//fmt.Println(s)
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(a)

	reverse(a[:])
	fmt.Println(a)

	rotate(a[:], 2)
	fmt.Println(a)

	var s1 []int //len(s1)=0, cap(s1)=0, s1 == nil? true
	fmt.Printf("len(s1)=%d, cap(s1)=%d, s1 == nil? %t\n", len(s1), cap(s1), s1 == nil)
	s1 = nil //len(s1)=0, cap(s1)=0, s1 == nil? true
	fmt.Printf("len(s1)=%d, cap(s1)=%d, s1 == nil? %t\n", len(s1), cap(s1), s1 == nil)
	s1 = []int(nil) //len(s1)=0, cap(s1)=0, s1 == nil? true
	fmt.Printf("len(s1)=%d, cap(s1)=%d, s1 == nil? %t\n", len(s1), cap(s1), s1 == nil)
	s1 = []int{} //len(s1)=0, cap(s1)=0, s1 == nil? false
	fmt.Printf("len(s1)=%d, cap(s1)=%d, s1 == nil? %t\n", len(s1), cap(s1), s1 == nil)

	s2 := make([]int, 3)
	fmt.Printf("s2 := make([]int, 3), len(s2)=%d, cap(s2)=%d, s2 == nil? %t\n", len(s2), cap(s2), s2 == nil)
	s3 := make([]int, 3, 6)
	fmt.Printf("s3 := make([]int, 3, 6), len(s3)=%d, cap(s3)=%d, s3 == nil? %t\n", len(s3), cap(s3), s3 == nil)
	s4 := make([]int, 6)[:3]
	fmt.Printf("s4 := make([]int, 6)[:3], len(s4)=%d, cap(s4)=%d, s4 == nil? %t\n", len(s4), cap(s4), s4 == nil)

}
