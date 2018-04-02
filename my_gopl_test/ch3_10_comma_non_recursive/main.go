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

	fmt.Println(comma("+12"))
	fmt.Println(comma("-123"))
	fmt.Println(comma("+1234"))
	fmt.Println(comma("-12345"))
	fmt.Println(comma("+123456"))
	fmt.Println(comma("-1234567"))

	fmt.Println(comma("+12.0"))
	fmt.Println(comma("-123.01"))
	fmt.Println(comma("+1234.012"))
	fmt.Println(comma("-12345.01234"))
	fmt.Println(comma("+123456.012345"))
	fmt.Println(comma("-1234567.0123456"))

}

//comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	var buff bytes.Buffer
	var rangeStartIndex, rangeEndIndex int

	if bytes.HasPrefix([]byte(s), []byte("+")) || bytes.HasPrefix([]byte(s), []byte("-")) {
		n -= 1
		buff.WriteByte(s[0])
		rangeStartIndex = 1
	}

	rangeEndIndex = len(s)
	dotIndex := bytes.IndexByte([]byte(s), '.')
	if dotIndex != -1 {
		n -= len(s) - dotIndex
		rangeEndIndex = dotIndex
	}

	if n <= 3 {
		return s
	}

	k := 3 - (n % 3)

	for i, c := range []byte(s[rangeStartIndex:rangeEndIndex]) {
		if i != 0 && (i+k)%3 == 0 {
			buff.WriteByte(',')
		}
		buff.WriteByte(c)
	}

	if dotIndex != -1 {
		buff.Write([]byte(s[dotIndex:]))
	}

	return buff.String()
}
