package main

import (
	"fmt"
    "github.com/jmptc/monkey-interpreter/lexer"
    "github.com/jmptc/monkey-interpreter/parser"
	"github.com/jmptc/monkey-interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

    // temp for testing
    input := "let x = 1 + 2;"

    if input != "" {
        l := lexer.New(input)
        p := parser.New(l)
        program := p.ParseProgram()
        fmt.Println(program.String())
    } else {
        fmt.Printf("Hello %s! This is the Monkey programming language!\n",
            user.Username)
        fmt.Printf("Feel free to type in commands\n")
        repl.Start(os.Stdin, os.Stdout)
    }
}
