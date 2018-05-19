package main

import (
	"fmt"
)

func max(val0 int, vals ...int) int {
	maxVal := val0
	for _, val := range vals {
		if maxVal < val {
			maxVal = val
		}
	}
	return maxVal
}

func min(val0 int, vals ...int) int {
	minVal := val0
	for _, val := range vals {
		if minVal > val {
			minVal = val
		}
	}
	return minVal
}

func main() {
	fmt.Println(max(0))
	fmt.Println(max(0, 1, 2, 3))
	fmt.Println(max(0, 1, 2, 3, -1, -2))

	fmt.Println(min(0))
	fmt.Println(min(0, 1, 2, 3))
	fmt.Println(min(0, 1, 2, 3, -1, -2))

}
