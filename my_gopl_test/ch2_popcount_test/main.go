package main

import (
	"fmt"
	"github.com/wangyoucao577/golang_test/my_gopl_test/ch2_popcount"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseUint(arg, 0, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ch2_popcount_test: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(popcount.PopCount(t))
	}

}
