// prints its command-line arguments
package main

import(
    "fmt"
    "os"
)

func main(){
    for index, arg := range os.Args[0:] {
        fmt.Println(index, " ", arg)
    }
}


