package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Point  // 匿名成员, 实际成员名就是类型名Point
	Radius int
}

type Wheel struct {
	Circle // 匿名成员, 实际成员名就是类型名Point
	Spokes int
}

func main() {
	w1 := Wheel{Circle{Point{8, 8}, 5}, 20}
	w2 := Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5, // `,` necessary here
		},
		Spokes: 20, // `,` necessary here
	}

	fmt.Printf("%#v\n", w1)
	fmt.Printf("%#v\n", w2)
}
