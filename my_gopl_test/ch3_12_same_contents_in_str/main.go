package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {

	fmt.Println(compareStringsEvenDisorder("", ""))
	fmt.Println(compareStringsEvenDisorder("a", "a"))
	fmt.Println(compareStringsEvenDisorder("ab", "ab"))
	fmt.Println(compareStringsEvenDisorder("ab", "ba"))
	fmt.Println(compareStringsEvenDisorder("abc", "acb"))
	fmt.Println(compareStringsEvenDisorder("abc", "abcd"))
	fmt.Println(compareStringsEvenDisorder("abcdee", "abcde"))
}

func compareStringsEvenDisorder(s1 string, s2 string) int {

	fmt.Printf("compareStringsEvenDisorder(%s, %s): ", s1, s2)

	b1 := []byte(s1)
	b2 := []byte(s2)

	sort.Slice(b1, func(i, j int) bool {
		return b1[i] > b1[j]
	})
	sort.Slice(b2, func(i, j int) bool {
		return b2[i] > b2[j]
	})

	return bytes.Compare(b1, b2)
}
