package parser

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

var VaslidPathChar = [...]rune{
	'M', 'm', 'L', 'l',
	'H', 'h', 'V', 'v',
	'C', 'c', 'S', 's',
	'Q', 'q', 'T', 't',
	'A', 'a', 'Z', 'z',
}

type PathMotion struct {
	Letter rune
	Values []float64
}

type PathMotions []PathMotion

func createNextMotion(char rune, motions *PathMotions) error {
	err := CheckCharIsValidMotionb(char)
	if err != nil {
		return err
	}

	*motions = append(*motions, PathMotion{char, make([]float64, 0, 7)})

	return nil
}

func incrementCurrentPathPoint(char rune, motions *PathMotions, currentMotionIndex, currentValuesIndex, charIncrement, decimalIndex int, isNegative bool) error {
	if currentMotionIndex < 0 {
		return errors.New("a letter is expected to begin a path")
	}
	if currentValuesIndex >= 7 {
		return fmt.Errorf("too many values for motion %c provided at pos %v", (*motions)[currentMotionIndex].Letter, charIncrement)
	}

	number, err := strconv.ParseFloat(string(char), 64)
	if err != nil {
		return fmt.Errorf("character %c at position %v could not be converted to a float2", char, charIncrement)
	}
	if isNegative {
		number *= -1
	}

	if len((*motions)[currentMotionIndex].Values) < currentValuesIndex+1 {
		(*motions)[currentMotionIndex].Values = append((*motions)[currentMotionIndex].Values, 0)
	}

	if decimalIndex > 0 {
		factor := 1 / math.Pow10(decimalIndex)
		(*motions)[currentMotionIndex].Values[currentValuesIndex] += number * factor
	} else {
		(*motions)[currentMotionIndex].Values[currentValuesIndex] *= 10
		(*motions)[currentMotionIndex].Values[currentValuesIndex] += number
	}
	return nil
}

func ParseMotions(fullPath string) (PathMotions, error) {
	var motions PathMotions
	currentMotionIndex := -1
	decimalIndex := 0
	currentValuesIndex := 0
	isParsingAdigit := false
	isNegative := false

	for pos, char := range fullPath {
		switch {

		case unicode.IsLetter(char):
			err := createNextMotion(char, &motions)

			if err != nil {
				return nil, err
			}

			currentMotionIndex++
			currentValuesIndex = 0
			isParsingAdigit = false
			decimalIndex = 0

		case unicode.IsDigit(char):
			isParsingAdigit = true
			err := incrementCurrentPathPoint(char, &motions, currentMotionIndex, currentValuesIndex, pos, decimalIndex, isNegative)

			if err != nil {
				return nil, err
			}

			if decimalIndex > 0 {
				decimalIndex++
			}

		case char == '.': // char is period
			decimalIndex++

		case char == '-':
			isNegative = true

		case char == ' ' || char == ',': // char is space or ,
			decimalIndex = 0
			isNegative = false

			if isParsingAdigit {
				currentValuesIndex++
			}

		default:
			return nil, fmt.Errorf("invalid character %c at position %v", char, pos)
		}
	}
	return motions, nil
}

// could be a binary search if array was sorted
func CheckCharIsValidMotionb(char rune) error {

	for _, a := range VaslidPathChar {
		if a == char {
			return nil
		}
	}

	return errors.New(string(char) + " is not a valid path character.")
}

func StringifyMotions(motions *PathMotions) string {
	var output strings.Builder

	for i, motion := range *motions {
		output.WriteString(string(motion.Letter))
		for _, val := range motion.Values {
			output.WriteString(strconv.FormatFloat(val, 'f', -1, 64))
			output.WriteString(string(','))
		}

		if i < len(*motions)-1 {
			output.WriteString(string(' '))
		}

	}
	return output.String()
}
