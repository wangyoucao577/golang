package main

import (
	"fmt"

	"github.com/wangyoucao577/golang_test/my_gopl_test/ch6_intset"
)

func main() {
	var x, y ch6_intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"
	fmt.Println(x.Len())    // 3

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"
	fmt.Println(y.Len())    // 2

	x.UnionWith(&y)
	fmt.Println(x.String())           // "{1 9 42 144}"
	fmt.Println(x.Len())              // 4
	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)

	z := x.Copy()
	fmt.Println(z) // "{1 9 42 144}"

	x.Remove(2)
	fmt.Println(&x) // "{1 9 42 144}"
	x.Remove(42)
	fmt.Println(&x) // "{1 9 144}"
	x.Remove(145)
	fmt.Println(&x) // "{1 9 144}"
	x.Remove(144)
	fmt.Println(&x) // "{1 9}"

	fmt.Println(x.Len()) // 2
	fmt.Println(&x)      // "{1 9}"
	x.Clear()
	fmt.Println(x.Len()) // 0
	fmt.Println(&x)      // "{}"

	fmt.Println(z) // "{1 9 42 144}"

	x.AddAll(1, 2, 5)
	fmt.Println(&x) // "{1, 2, 5}"
	x.AddAll(4, 10)
	fmt.Println(&x) // "{4, 10}"
}
