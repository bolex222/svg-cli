package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bolex222/svg-cli/internal/flagmanagement"
	// "github.com/bolex222/svg-cli/internal/parser"
	"github.com/bolex222/svg-cli/internal/tokenizer"
)

func main() {
	options := flagmanagement.ParseFlags()
	path, err := flagmanagement.GetPath()

	if err != nil {
		fmt.Printf("%v \n", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(strings.NewReader(path))

	tok := tokenizer.New()
	tokens, err := tok.Tokenize(reader)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i, t := range tokens {
		fmt.Println(t, i)
	}

	// TODO: tokenize from ioReader

	// TODO: parse from tokens

	for _, opt := range options {
		fmt.Printf("opt %v has value %v \n", opt.Name, opt.Value)
	}

	// for _, motion := range  parsedPath {
	// 	fmt.Printf("motion %c \n", motion.Letter)
	// }
}
