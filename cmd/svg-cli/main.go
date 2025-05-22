package main

import (
	"fmt"
	"os"

	"github.com/bolex222/svg-cli/internal/parser"
)

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Println("The path are required")
		os.Exit(1)
	}

	allMotionsAsString := args[1]
	motions, err := parser.ParseMotions(allMotionsAsString)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, val := range motions {
		fmt.Printf("motion: %c, values: %v\n", val.Letter, val.Values)
	}

}
