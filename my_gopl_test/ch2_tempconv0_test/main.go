// main test for tempconv0
package main

import (
	"fmt"
	"github.com/wangyoucao577/golang_test/my_gopl_test/ch2_tempconv0"
)

func main() {
	fmt.Printf("%g\n", tempconv0.BoilingC-tempconv0.FreezingC) // "100"
	boilingF := tempconv0.CToF(tempconv0.BoilingC)
	fmt.Printf("%g\n", boilingF-tempconv0.CToF(tempconv0.FreezingC)) // "180"
	//fmt.Printf("%g\n", boilingF-FreezingC)       // compile error: type mismatch

	var c tempconv0.Celsius
	var f tempconv0.Fahrenheit
	fmt.Println(c == 0) // "true"
	fmt.Println(f >= 0) // "true"
	//fmt.Println(c == f)    // compile error: type mismatch
	fmt.Println(c == tempconv0.Celsius(f)) // "true"

	c2 := tempconv0.FToC(212.0)
	fmt.Println(c2.String()) // "100C"
	fmt.Printf("%v\n", c2)   // "100C"; no need to call String() explicity
	fmt.Printf("%s\n", c2)   // "100C"
	fmt.Println(c2)          // "100C"
	fmt.Printf("%g\n", c2)   // "100"; does not call String()
	fmt.Println(float64(c2)) // "100"; does not call String()

}
