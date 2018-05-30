package main

import (
	"fmt"
	"io"
	"os"
)

type countingWriter struct {
	counter *int64
	writer  io.Writer
}

func (c countingWriter) Write(p []byte) (int, error) {
	*(c.counter) += int64(len(p))
	return c.writer.Write(p)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var c int64
	cw := countingWriter{&c, w}
	return cw, cw.counter
}

func main() {
	cw, count := CountingWriter(os.Stdout)
	fmt.Fprintf(cw, "hello world!\n")
	//fmt.Println(cw)
	fmt.Println(*count)
}
