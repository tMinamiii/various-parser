package main

import (
	"fmt"
	"os"
	"os/user"

	monkey "github.com/tMinamiii/various-parser/monkey/parser"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	monkey.StartREPL(os.Stdin, os.Stdout)
}
