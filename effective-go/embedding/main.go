package main

import (
	"fmt"
	"io"
	"reflect"
)

type reader int

func (r reader) Read(p []byte) (n int, err error) {
	//TODO: do read
	return len(p), nil
}

type readerNoFieldName struct {
	reader
}

type readerWithFieldName struct {
	reader reader
}

var _ io.Reader = readerNoFieldName{}

//var _ io.Reader = readerWithFieldName{}	// Can not compile

func satisfyReaderCheck(name string, v interface{}) {
	if _, ok := v.(io.Reader); ok {
		fmt.Printf("type %s satisfied io.Reader\n", name)
		return
	}
	fmt.Printf("type %s didn't satisfy io.Reader\n", name)
}

func main() {

	r1 := readerNoFieldName{}
	satisfyReaderCheck(reflect.TypeOf(r1).String(), r1)

	r2 := readerWithFieldName{}
	satisfyReaderCheck(reflect.TypeOf(r2).String(), r2)

}
