package main

import (
	"fmt"
)

func returnNonZeroButNoReturnSentence() (val int) {
	type myExpectedPanic struct{}
	var myPanic myExpectedPanic

	defer func() {
		if p := recover(); p == myPanic {
			val = 1
		}
	}()

	panic(myPanic)
}

func main() {
	fmt.Println(returnNonZeroButNoReturnSentence())
}
