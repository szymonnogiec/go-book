package main

import (
	"fmt"
	"os"
)

func main() {
	for i, s := range os.Args[1:] {
		fmt.Println(s)
		fmt.Println(i)
	}
}
