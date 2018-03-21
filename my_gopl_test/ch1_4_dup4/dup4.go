package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	files := os.Args[1:]
	if len(files) == 0 {
		counts := make(map[string]int)
		countLines(os.Stdin, counts)
		printDup("os.Stdin", counts)
	} else {
		for _, arg := range files {
			counts := make(map[string]int)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup4: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()

			printDup(arg, counts)
		}
	}

}

func printDup(filename string, counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\t%s\n", filename, n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
