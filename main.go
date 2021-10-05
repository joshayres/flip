package main

import (
	"flip/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s, welcome to flip!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
