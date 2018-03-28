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
}
