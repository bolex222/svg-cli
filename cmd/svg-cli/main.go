package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bolex222/svg-cli/internal/flagmanagement"
	"github.com/bolex222/svg-cli/internal/parser"
	"github.com/bolex222/svg-cli/internal/tokenizer"
)

func main() {
	flagmanagement.ParseFlags()
	path, err := flagmanagement.GetPath()

	if err != nil {
		fmt.Printf("%v \n", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(strings.NewReader(path))

	tok := tokenizer.Tokenizer{}

	tokens, err := tok.Tokenize(reader)

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

	for i, t := range commands {
		fmt.Printf("index --> %v : command --> %c \n", i, t.Letter)
	}
}
