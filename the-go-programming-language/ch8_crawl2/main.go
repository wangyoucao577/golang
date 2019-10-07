package main

import (
	"fmt"
	"log"
	"os"

	links "github.com/wangyoucao577/golang_test/my_gopl_test/ch5_links"
)

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent request.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} //acquire a token

	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	<-tokens // release the token
	return list
}

func main() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	//Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	//Crawl the web concurrently.
	seen := make(map[string]bool)

	for ; n > 0; n-- { //wait for all goroutine returns
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++

				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}

	}
}
