package main

import (
	"fmt"

	"github.com/wangyoucao577/golang_test/my_gopl_test/ch6_geometry"
)

func main() {
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}

	fmt.Println(geometry.Distance(p, q))
	fmt.Println(p.Distance(q))

	path1 := geometry.Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	fmt.Println(path1.Distance())
}
