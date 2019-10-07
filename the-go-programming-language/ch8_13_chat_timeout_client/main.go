package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			//Broadcast incoming message to call
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}

	}
}

const MaxSecondsToLive = 20

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	tick := time.NewTicker(1 * time.Second)
	defer tick.Stop()

	inMsg := make(chan string)
	go clientReader(conn, who, inMsg)

	countdown := MaxSecondsToLive
loop:
	for {
		select {
		case msg := <-inMsg:
			countdown = MaxSecondsToLive
			messages <- msg
		case <-tick.C:
			countdown--
			fmt.Printf("Client %s remain %d seconds to live.\n", who, countdown)
			if countdown <= 0 {
				break loop
			}
		}
	}

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientReader(conn net.Conn, who string, in chan<- string) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		in <- who + ": " + input.Text()
	}
	//NOTE: ignoring potential erros from input.Err()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}
