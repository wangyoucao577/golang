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
	fmt.Printf("Connection %v Enter Echo, text %s\n", c.RemoteAddr(), shout)

	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))

	fmt.Printf("Connection %v Exit Echo, text %s\n", c.RemoteAddr(), shout)

	wg.Done()
}

const (
	ExpireSeconds = 10
)

func handleConn(c net.Conn) {
	fmt.Printf("Connection %v Entered\n", c.RemoteAddr())

	defer c.Close()

	var wg sync.WaitGroup

	incoming := make(chan string)
	go func(c net.Conn) {
		input := bufio.NewScanner(c)
		for input.Scan() { //NOTE: ignoring potenial errors from input.Err()
			incoming <- input.Text()
		}
	}(c)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	countdown := ExpireSeconds
	for {
		select {
		case text := <-incoming:
			wg.Add(1)
			countdown = ExpireSeconds // ignore delay in echo
			go echo(c, text, 1*time.Second, &wg)
		case <-ticker.C:
			fmt.Printf("Connection %v Remain Seconds %d\n", c.RemoteAddr(), countdown)
			countdown--
			if countdown <= 0 {
				fmt.Printf("Connection %v Exit\n", c.RemoteAddr())
				return // exit connection
			}
		}
	}
}
