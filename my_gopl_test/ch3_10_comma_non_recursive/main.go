package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("12"))
	fmt.Println(comma("123"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))
}

//comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	k := 3 - (n % 3)
	var buff bytes.Buffer
	for i, c := range []byte(s) {
		if i != 0 && (i+k)%3 == 0 {
			buff.WriteByte(',')
		}
		buff.WriteByte(c)
	}

	return buff.String()
}
