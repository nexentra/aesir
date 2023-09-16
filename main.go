package main

import (
	"fmt"
	"github.com/nexentra/aesir/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	
	banner()
	fmt.Printf("Hello %s! This is the AEsir programming language!\n",user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

func banner() {
	b, err := os.ReadFile("ascii_art.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}