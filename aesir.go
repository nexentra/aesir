package aesir

import (
	"github.com/nexentra/aesir/evaluator"
	"github.com/nexentra/aesir/lexer"
	"github.com/nexentra/aesir/object"
	"github.com/nexentra/aesir/parser"
)

func Execute(input string, env *object.Environment) (string, []string) {
	l := lexer.New(input)
	p := parser.New(l)
	if len(p.Errors()) != 0 {
		return "", p.Errors()
	}
	program := p.ParseProgram()
	evaluated := evaluator.Eval(program, env)
	return evaluated.Inspect(), nil
}