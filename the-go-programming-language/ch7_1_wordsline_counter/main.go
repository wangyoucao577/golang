package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c += 1
	}
	return len(p), nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*c += 1
	}
	return len(p), nil
}

func main() {
	var c1 WordCounter
	fmt.Println(c1)
	fmt.Fprintf(&c1, "hello world\nhello golang!")
	fmt.Println(c1)

	var c2 LineCounter
	fmt.Println(c2)
	fmt.Fprintf(&c2, "hello world\nhello golang!")
	fmt.Println(c2)

}
