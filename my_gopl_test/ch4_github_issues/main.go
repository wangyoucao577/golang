// Issues prints a table of Github issues matching the search terms.
package main

import (
	"fmt"
	"github.com/wangyoucao577/golang_test/my_gopl_test/ch4_github"
	"log"
	"os"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
