package main

import (
	"fmt"
	"strings"
)

func join(sep string, s ...string) string {
	return strings.Join(s, sep)
}

func main() {
	fmt.Println(join(", "))
	fmt.Println(join(", ", "123"))
	fmt.Println(join(", ", "123", "456", "789"))
}
