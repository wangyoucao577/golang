package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	fmt.Printf("%T\n", w) //"<nil>"
	fmt.Println(w)
	w = os.Stdout
	fmt.Printf("%T\n", w) //"*os.File"
	fmt.Println(w)
	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w) //"*bytes.Buffer"
	fmt.Println(w)
	w = nil
	fmt.Printf("%T\n", w)
	fmt.Println(w)
}
