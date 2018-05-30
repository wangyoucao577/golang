package main

import (
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter((len(p)))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // 5

	c = 0
	name := "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // 12
}
