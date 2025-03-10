package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/claudio4/monkey-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to this toy interpreter for the toy language Monkey!\n", user.Name)
	repl.Start(os.Stdin, os.Stdout)
}
