package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle one connection at a time
	}
}

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	log.Printf("Enter new echo routine, remote addr: %s\n", c.RemoteAddr())

	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))

	log.Printf("Exit new echo routine, remote addr: %s\n", c.RemoteAddr())
}

func handleConn(c net.Conn) {
	log.Printf("Handle new connection, remote addr: %s\n", c.RemoteAddr())

	var wg sync.WaitGroup
	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		go echo(c, input.Text(), 5*time.Second, &wg)
	}

	log.Printf("Before wait group, remote addr: %s\n", c.RemoteAddr())
	wg.Wait()

	log.Printf("Ready to close connection, remote addr: %s\n", c.RemoteAddr())
	//NOTE: ignoring potenial errors from input.Err()
	c.Close()
}
