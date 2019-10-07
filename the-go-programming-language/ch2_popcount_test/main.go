package main

import (
	"flag"
	"fmt"
	"github.com/wangyoucao577/golang_test/my_gopl_test/ch2_popcount"
	"time"
)

var loopCount = flag.Int("c", 1, "loop count")
var value = flag.Uint64("v", 0x1234567890ABCDEF, "test value")

func main() {
	flag.Parse()
	t := *value

	var p1, p2, p3, p4 int

	start := time.Now()
	for i := 0; i < *loopCount; i++ {
		p1 = popcount.PopCountByLookupTable(t)
	}
	end := time.Now()
	fmt.Printf("PopCountByLookupTable : %d cost %d ms (loop count %d)\n", p1, end.Sub(start)/time.Millisecond, *loopCount)

	start = time.Now()
	for i := 0; i < *loopCount; i++ {
		p2 = popcount.PopCountByLookupTable(t)
	}
	end = time.Now()
	fmt.Printf("PopCountByLooping     : %d cost %d ms (loop count %d)\n", p2, end.Sub(start)/time.Millisecond, *loopCount)

	start = time.Now()
	for i := 0; i < *loopCount; i++ {
		p3 = popcount.PopCountByShifting(t)
	}
	end = time.Now()
	fmt.Printf("PopCountByShifting    : %d cost %d ms (loop count %d)\n", p3, end.Sub(start)/time.Millisecond, *loopCount)

	start = time.Now()
	for i := 0; i < *loopCount; i++ {
		p4 = popcount.PopCountByClearing(t)
	}
	end = time.Now()
	fmt.Printf("PopCountByClearing    : %d cost %d ms (loop count %d)\n", p4, end.Sub(start)/time.Millisecond, *loopCount)

}
