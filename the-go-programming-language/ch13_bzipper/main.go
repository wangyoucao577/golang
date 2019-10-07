// Bzipper reads input, bzip2-compresses it, and writes it out.
package main

import (
	"io"
	"log"
	"os"

	bzip "github.com/wangyoucao577/golang_test/my_gopl_test/ch13_bzip"
)

func main() {
	w := bzip.NewWriter(os.Stdout)
	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatalf("bzipper: %v\n", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("bzipper: close: %v", err)
	}

}
