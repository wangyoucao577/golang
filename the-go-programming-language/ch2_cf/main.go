// Cf converts its numeric argument to Celsius and Fahrenheit
package main

import (
	"fmt"
	"github.com/wangyoucao577/golang_test/my_gopl_test/ch2_tempconv0"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv0.Fahrenheit(t)
		c := tempconv0.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv0.FToC(f), c, tempconv0.CToF(c))
	}
}
