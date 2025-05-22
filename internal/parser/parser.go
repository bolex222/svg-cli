package parser

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"strconv"
	"unicode"
)

var VaslidPathChar = []rune{
	rune('M'), rune('m'), rune('L'), rune('l'),
	rune('H'), rune('h'), rune('V'), rune('v'),
	rune('C'), rune('c'), rune('S'), rune('s'),
	rune('Q'), rune('q'), rune('T'), rune('t'),
	rune('A'), rune('a'), rune('Z'), rune('z'),
}

type pathMotion struct {
	Letter rune
	Values [7]float64
}

type pathMotions []pathMotion

func createNextMotion(char rune, motions *pathMotions) error {
	err := CheckCharIsValidMotionb(char)
	if err != nil {
		return err
	}

	*motions = append(*motions, pathMotion{
		Letter: char,
	})

	return nil
}

func incrementCurrentPathPoint(char rune, motions *pathMotions, currentMotionIndex, currentValuesIndex, charIncrement, decimalIndex int) error {

	if currentMotionIndex < 0 {
		return errors.New("a letter is expected to begin a path")
	}

	if currentValuesIndex >= 7 {
		return fmt.Errorf("too many values for motion %c provided at pos %v", (*motions)[currentMotionIndex].Letter, charIncrement)
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

func ParseMotions(fullPath string) ([]pathMotion, error) {
	var motions pathMotions
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

func CheckCharIsValidMotionb(char rune) error {
	if !slices.Contains(VaslidPathChar, char) {
		return errors.New(string(char) + " is not a valid path character.")
	}

	return nil
}
