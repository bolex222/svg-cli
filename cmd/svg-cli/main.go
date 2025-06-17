package main

import (
	"fmt"
	"os"

	"github.com/bolex222/svg-cli/internal/flagmanagment"
)

func main() {
	options := flagmanagment.ParseFlags()
	path, err := flagmanagment.GetPath()

	if err != nil {
		fmt.Printf("%v \n", err)
		os.Exit(1)
	}
	
	fmt.Printf("path: %v \n", path )


	for _, opt := range options {
		fmt.Printf("opt %v has value %v \n", opt.Name, opt.Value)
	}
}
