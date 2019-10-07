package main

import (
	"fmt"
	"log"
	"os"

	links "github.com/wangyoucao577/golang_test/my_gopl_test/ch5_links"
)

type linkWithMetadata struct {
	Url   string
	Depth int
}

const (
	MaxDepth = 3
)

var tokens = make(chan struct{}, 20) //max 20 concurrent

func crawl(url string) []string {
	//fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func urlsTolinkWithMetadatas(urls []string, depth int) []linkWithMetadata {
	var links []linkWithMetadata
	for _, url := range urls {
		links = append(links, linkWithMetadata{url, depth})
	}
	return links
}

func main() {
	worklist := make(chan []linkWithMetadata)
	validUnseen := make(chan linkWithMetadata)
	count := make(chan int)

	go func() {
		worklist <- urlsTolinkWithMetadatas(os.Args[1:], 1) //input urls
	}()

	go func() {
		var n int
		for {
			n += <-count
			fmt.Printf("totally seen %d\n", n)
		}
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for unseen := range validUnseen {
				fmt.Printf("[%d]%s\n", unseen.Depth, unseen.Url)
				urls := crawl(unseen.Url)

				nextDepth := unseen.Depth + 1
				if nextDepth <= MaxDepth {
					go func(nextDepth int) {
						worklist <- urlsTolinkWithMetadatas(urls, nextDepth)
					}(nextDepth)
				}
				count <- 1
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, item := range list {
			if item.Depth <= MaxDepth && !seen[item.Url] {
				seen[item.Url] = true

				validUnseen <- item
			}
		}

		//TODO: How to exit this loop?
	}

}
