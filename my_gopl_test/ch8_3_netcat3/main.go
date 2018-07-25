package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) //NOTE: ignoring error
		log.Println("done")
		done <- struct{}{}
	}()
	if tcpConn, ok := conn.(*net.TCPConn); ok {
		tcpConn.CloseWrite()
		time.Sleep(10 * time.Second)
	}

	conn.Close()
	log.Println("wait for background goroutine to finish")
	<-done // wait for background goroutine to finish
	log.Println("wait done")
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
