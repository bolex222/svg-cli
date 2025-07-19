package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bolex222/svg-cli/internal/command"
	"github.com/bolex222/svg-cli/internal/flagmanagement"
	"github.com/bolex222/svg-cli/internal/lexer"
	"github.com/bolex222/svg-cli/internal/parser"
)

func main() {
	flagmanagement.ParseFlags()
	path, err := flagmanagement.GetPath()

	if err != nil {
		fmt.Printf("%v \n", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(strings.NewReader(path))
	lex := lexer.New()
	tokens, err := lex.Tokenize(reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	par := parser.New()
	commands, err := par.ParseTokensToCommands(tokens)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("amount of command --> %v \n", len(commands))

	for _, t := range commands {
		fmt.Printf("/////////// Command --> %c ///////////// \n", t.Letter)
		fmt.Printf("values --> %v \n", t.Values)
		if t.Type == command.ElipticArcValueCommand {
			fmt.Printf("Angle --> %v \n", t.Angle)
			fmt.Printf("LargeArcFlag --> %v \n", t.LargeArcFlag)
			fmt.Printf("SweepFlag --> %v \n", t.SweepFlag)
		}
	}
}
