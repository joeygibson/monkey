package main

import (
	"fmt"
	"os"
	osUser "os/user"

	"github.com/joeygibson/monkey/repl"
)

func main() {
	user, err := osUser.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello, %s! Welcome to Monkey!\n", user.Username)
	fmt.Printf("Type some commands.\n")

	repl.Start(os.Stdin, os.Stdout)
}
