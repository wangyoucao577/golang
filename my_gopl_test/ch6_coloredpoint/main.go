package main

import (
	"fmt"
	"image/color"

	"github.com/wangyoucao577/golang_test/my_gopl_test/ch6_geometry"
)

type ColoredPoint struct {
	geometry.Point
	Color color.RGBA
}

func main() {
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}

	fmt.Println(geometry.Distance(p, q))
	fmt.Println(p.Distance(q))

	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) //1
	cp.Point.Y = 2
	fmt.Println(cp.Y) //2

	fmt.Println(cp.Distance(q))

}
