package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("hello, %s!, this is mokey language\n", user.Username)
	fmt.Printf("Free free to type in command\n")

	repl.Start(os.Stdin, os.Stdout)
}
