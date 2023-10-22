package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/user"

	"github.com/nexentra/aesir/importer"
)

const AESIR = `

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

	args := os.Args[1:]

	if len(args) != 0 {
		if _, err := os.Stat(args[0]); err == nil {
			_, err := importer.Importer(string(args[0]))
			if err != nil {
				panic(err)
			}
		} else if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File does not exist")

		}

	} else {
		io.WriteString(os.Stdout, AESIR)
		fmt.Printf("Hello %s! This is the AEsir programming language!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		Start(os.Stdin, os.Stdout)
	}
}
