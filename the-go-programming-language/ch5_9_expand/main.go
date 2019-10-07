package main

import (
	"fmt"
	"os"
	"strings"
)

func f1(s string) string {
	return strings.ToUpper(s)
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "foo", f("foo"), -1)
}

func main() {

	for _, s := range os.Args[1:] {
		fmt.Printf("%s --> %s\n", s, expand(s, f1))
	}

}
