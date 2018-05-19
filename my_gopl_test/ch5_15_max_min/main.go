package main

import (
	"fmt"
)

func max(vals ...int) int {
	if len(vals) == 0 {
		fmt.Println("At least one param should be passed in.")
		return 0
	}

	maxVal := vals[0]
	for _, val := range vals {
		if maxVal < val {
			maxVal = val
		}
	}
	return maxVal
}

func min(vals ...int) int {
	if len(vals) == 0 {
		fmt.Println("At least one param should be passed in.")
		return 0
	}

	minVal := vals[0]
	for _, val := range vals {
		if minVal > val {
			minVal = val
		}
	}
	return minVal
}

func main() {
	fmt.Println(max())
	fmt.Println(max(0))
	fmt.Println(max(0, 1, 2, 3))
	fmt.Println(max(0, 1, 2, 3, -1, -2))

	fmt.Println(min())
	fmt.Println(min(0))
	fmt.Println(min(0, 1, 2, 3))
	fmt.Println(min(0, 1, 2, 3, -1, -2))

}
