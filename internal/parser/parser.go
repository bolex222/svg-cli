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
	rune('M'), rune('m'), rune('L'), rune('l'),
	rune('H'), rune('h'), rune('V'), rune('v'),
	rune('C'), rune('c'), rune('S'), rune('s'),
	rune('Q'), rune('q'), rune('T'), rune('t'),
	rune('A'), rune('a'), rune('Z'), rune('z'),
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

func incrementCurrentPathPoint(char rune, motions *PathMotions, currentMotionIndex, currentValuesIndex, charIncrement, decimalIndex int) error {

	if currentMotionIndex < 0 {
		return errors.New("a letter is expected to begin a path")
	}

	if currentValuesIndex >= 7 {
		return fmt.Errorf("too many values for motion %c provided at pos %v", (*motions)[currentMotionIndex].Letter, charIncrement)
	}

	if len((*motions)[currentMotionIndex].Values) < currentValuesIndex+1 {
		(*motions)[currentMotionIndex].Values = append((*motions)[currentMotionIndex].Values, 0)
	}

	if decimalIndex > 0 {
		factor := 1 / math.Pow10(decimalIndex)
		parsedChar, err := strconv.ParseFloat(string(char), 64)

		if err != nil {
			return fmt.Errorf("character %c at position %v could not be converted to a float2", char, charIncrement)
		}

		(*motions)[currentMotionIndex].Values[currentValuesIndex] += parsedChar * factor
	} else {
		newVal, err := strconv.ParseFloat(string(char), 64)

		if err != nil {
			return fmt.Errorf("character %c at position %v could not be converted to a float2", char, charIncrement)
		}

		(*motions)[currentMotionIndex].Values[currentValuesIndex] *= 10
		(*motions)[currentMotionIndex].Values[currentValuesIndex] += newVal
	}
	return nil
}

func ParseMotions(fullPath string) (PathMotions, error) {
	var motions PathMotions
	charIncrement := 0
	currentMotionIndex := -1
	decimalIndex := 0
	currentValuesIndex := 0
	currentDigitIndex := 0

	for pos, char := range fullPath {
		switch {

		case unicode.IsLetter(char):
			err := createNextMotion(char, &motions)

			if err != nil {
				return nil, err
			}

			currentMotionIndex++
			currentValuesIndex = 0
			currentDigitIndex = 0
			decimalIndex = 0

		case unicode.IsDigit(char):
			currentDigitIndex++
			err := incrementCurrentPathPoint(char, &motions, currentMotionIndex, currentValuesIndex, charIncrement, decimalIndex)

			if err != nil {
				return nil, err
			}

			if decimalIndex > 0 {
				decimalIndex++
			}

		case char == 46: // char is period
			decimalIndex++

		case char == 44 || char == 32: // char is space or ,
			if currentDigitIndex > 0 {
				currentValuesIndex++
				decimalIndex = 0
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
