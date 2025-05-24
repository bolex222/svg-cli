package main

import (
	"fmt"
	"os"

	"github.com/bolex222/svg-cli/internal/parser"
)

var table map[rune][]int = map[rune][]int{
	rune('M'): {0, 1},
	rune('m'): {0, 1},
	rune('L'): {0, 1},
	rune('l'): {0, 1},
	rune('h'): {0},
	rune('H'): {0},
	rune('V'): {0},
	rune('v'): {0},
	rune('C'): {0, 1, 2, 3, 4, 5},
	rune('c'): {0, 1, 2, 3, 4, 5},
	rune('S'): {0, 1, 2, 3},
	rune('s'): {0, 1, 2, 3},
	rune('q'): {0, 1, 2, 3},
	rune('Q'): {0, 1, 2, 3},
	rune('T'): {0, 1},
	rune('t'): {0, 1},
	rune('a'): {0, 1, 5, 6},
	rune('A'): {0, 1, 5, 6},
	rune('z'): {},
	rune('Z'): {},
}

func getHighestValueInMotions(motions *parser.PathMotions) (float64, error) {
	var highest float64 = -1

	for i, motion := range *motions {
		for j, motionValueIndex := range table[motion.Letter] {
			if motionValueIndex > 6 {
				return 0, fmt.Errorf("trying to reach an out of band value of Motion %v: %c", i+1, motion.Letter)
			}

			value := motion.Values[motionValueIndex]
			if i == 0 && j == 0 {
				highest = value
			} else if value > highest {
				highest = value
			}
		}
	}

	return highest, nil
}

func normalizeAllMotions(motions parser.PathMotions, maxValue float64) (parser.PathMotions, error) {
	factor := 1 / maxValue

	for i, motion := range motions {
		for _, motionValueIndex := range table[motion.Letter] {
			if motionValueIndex > 6 {
				return nil, fmt.Errorf("trying to reach an out of band value of Motion %v: %c", i+1, motion.Letter)
			}
			motions[i].Values[motionValueIndex] *= factor
		}
	}

	return motions, nil
}

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

	highest, err := getHighestValueInMotions(&motions)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	normalizedMotions, err := normalizeAllMotions(motions, highest)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := parser.StringifyMotions(&normalizedMotions)

	fmt.Println(output)

	// for _, val := range motions {
	// 	fmt.Printf("motion: %c, values: %v\n", val.Letter, val.Values)
	// }

	// fmt.Printf("/////////////////////////////////// \n")

	// for _, val := range normalizedMotions {
	// 	fmt.Printf("motion: %c, values: %v\n", val.Letter, val.Values)
	// }
}
