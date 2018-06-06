package main

import (
	"flag"
	"fmt"

	"github.com/wangyoucao577/golang_test/my_gopl_test/ch2_tempconv0"
)

// *celsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct{ tempconv0.Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "℃":
		f.Celsius = tempconv0.Celsius(value)
		return nil
	case "F", "℉":
		f.Celsius = tempconv0.FToC(tempconv0.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument mush have a quantity and a unit, e.g., "100C"
func CelsiusFlag(name string, value tempconv0.Celsius, usage string) *tempconv0.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
