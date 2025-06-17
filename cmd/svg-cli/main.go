package main

import (
	"fmt"
	"os"

	"github.com/bolex222/svg-cli/internal/flagmanagment"
)

func main() {
	options, err := flagmanagment.InitFalgs()
	if err != nil {
		os.Exit(1)
	}

	for _, opt := range *options {
		fmt.Printf(" opt %v has value %v \n", opt.Name, opt.Value)
	}
}
