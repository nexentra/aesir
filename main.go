package main

import (
	"fmt"
	"io"
	"os"
	"os/user"

	"github.com/nexentra/aesir/repl"
)

const AESIR = 
`

________   _______    ________   ___   ________     
|\   __  \ |\  ___ \  |\   ____\ |\  \ |\   __  \    
\ \  \|\  \\ \   __/| \ \  \___|_\ \  \\ \  \|\  \   
 \ \   __  \\ \  \_|/__\ \_____  \\ \  \\ \   _  _\  
  \ \  \ \  \\ \  \_|\ \\|____|\  \\ \  \\ \  \\  \| 
   \ \__\ \__\\ \_______\ ____\_\  \\ \__\\ \__\\ _\ 
    \|__|\|__| \|_______||\_________\\|__| \|__|\|__|
                         \|_________|                


`

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	
	io.WriteString(os.Stdout, AESIR)
	fmt.Printf("Hello %s! This is the AEsir programming language!\n",user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}