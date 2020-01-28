package main

import (
	"fmt"
	flags "github.com/jessevdk/go-flags"
)

var myOpts struct {
	Name    string `short:"n" long:"name" default:"World" description:"A name to say hello to."`
	Spanish bool   `short:"s" long:"spanish" description:"Use Spanish lang"`
}

func main() {
	flags.Parse(&myOpts)
	if myOpts.Spanish == true {
		fmt.Printf("Hola %s!\n", myOpts.Name)
	} else {
		fmt.Printf("Hello %s!\n", myOpts.Name)
	}
}
