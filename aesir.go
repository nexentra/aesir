package aesir

import (
	"fmt"

	"github.com/nexentra/aesir/evaluator"
	"github.com/nexentra/aesir/lexer"
	"github.com/nexentra/aesir/object"
	"github.com/nexentra/aesir/parser"
)

func Execute(input string, env *object.Environment) (string, []string) {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		// return "", p.Errors()
		fmt.Println(p.Errors())
	}
	evaluated := evaluator.Eval(program, env)
	return evaluated.Inspect(), nil
}