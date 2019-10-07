// converts its numeric argument to Meters and Miles
package main

import (
	"fmt"
	"github.com/wangyoucao577/golang_test/my_gopl_test/ch2_2_meters_miles_conv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			os.Exit(1)
		}
		miles := meters_miles_conv.Miles(t)
		meters := meters_miles_conv.Meters(t)
		fmt.Printf("%s = %s, %s = %s\n",
			miles, meters_miles_conv.MilesToMeters(miles), meters, meters_miles_conv.MetersToMiles(meters))
	}
}
