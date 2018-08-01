package main

import (
	"fmt"
	"log"
	"os"

	links "github.com/wangyoucao577/golang_test/my_gopl_test/ch5_links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string) // de-duplicated urls

	//Start with the command-line arguments.
	go func() { worklist <- os.Args[1:] }()

	//Create 20 crawler goroutines to fetch each url
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	//The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}

		//TODO: same issue: can not exit the loop?
	}
}
