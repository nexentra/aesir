package importer

import (
	"os"
	"path/filepath"
	"strings"

	"fmt"

	"github.com/nexentra/aesir/evaluator"
	"github.com/nexentra/aesir/lexer"
	"github.com/nexentra/aesir/object"
	"github.com/nexentra/aesir/parser"
)

// The suffix of a aesir file.
const suffix = ".ae"

func Importer(dir string) ([]string, error) {
	// var generatedFiles []string
	env := object.NewEnvironment()
	// var results []string 

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), suffix) {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			translatedContent := string(data)

			l := lexer.New(translatedContent)
			p := parser.New(l)
			program := p.ParseProgram()
			if len(p.Errors()) != 0 {
				fmt.Println(p.Errors())
			}
			evaluated := evaluator.Eval(program, env)
			fmt.Println(evaluated.Inspect())
		}
		return nil
	})

	// return generatedFiles, err
	return nil, err
}