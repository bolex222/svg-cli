package parser

import (
	"fmt"
	"unicode"

	"github.com/bolex222/svg-cli/internal/motion"
)

var VaslidPathChar = [...]rune{
	'M', 'm', 'L', 'l',
	'H', 'h', 'V', 'v',
	'C', 'c', 'S', 's',
	'Q', 'q', 'T', 't',
	'A', 'a', 'Z', 'z',
}

type Path []motion.Motion

func CheckCharIsValidMotionb(char rune) bool {

	for _, a := range VaslidPathChar {
		if a == char {
			return true
		}
	}

	return false
}

type Progress struct {
	Move        rune
	motionIndex int
}

/**
* parser is a global function using dependency injection to parse based on the right motion
 */
func ParsePathString(fullPath string) (Path, error) {
	finalPath := make(Path, 0, 1)
	progress := Progress{
		motionIndex: -1,
	}

	for i, char := range fullPath {

		switch {
		case unicode.IsLetter(char):
			motionBuffer, err := motion.InitMotion(char)
			if err != nil {
				return finalPath, fmt.Errorf("unexpected character %c at position %v", char, i)
			}

			finalPath = append(finalPath, motionBuffer)
			progress.motionIndex++
		case unicode.IsSpace(char):
		default:
			return finalPath, fmt.Errorf("unexpected character %c at position %v", char, i)

		}
	}

	return finalPath, nil
}

// func createNextMotion(char rune, motions *Path) error {
// 	err := CheckCharIsValidMotionb(char)
// 	if err != nil {
// 		return err
// 	}

// 	*motions = append(*motions, PathMotion{char, make([]float64, 0, 7)})

// 	return nil
// }

// func incrementCurrentPathPoint(char rune, motions *Path, currentMotionIndex, currentValuesIndex, charIncrement, decimalIndex int, isNegative bool) error {
// 	if currentMotionIndex < 0 {
// 		return errors.New("a letter is expected to begin a path")
// 	}
// 	if currentValuesIndex >= cap((*motions)[currentMotionIndex].Values) {
// 		return fmt.Errorf("too many values for motion %c provided at pos %v", (*motions)[currentMotionIndex].Letter, charIncrement)
// 	}

// 	number, err := strconv.ParseFloat(string(char), 64)
// 	if err != nil {
// 		return fmt.Errorf("character %c at position %v could not be converted to a float2", char, charIncrement)
// 	}
// 	if isNegative {
// 		number *= -1
// 	}

// 	if len((*motions)[currentMotionIndex].Values) < currentValuesIndex+1 {
// 		(*motions)[currentMotionIndex].Values = append((*motions)[currentMotionIndex].Values, 0)
// 	}

// 	if decimalIndex > 0 {
// 		factor := 1 / math.Pow10(decimalIndex)
// 		(*motions)[currentMotionIndex].Values[currentValuesIndex] += number * factor
// 	} else {
// 		(*motions)[currentMotionIndex].Values[currentValuesIndex] *= 10
// 		(*motions)[currentMotionIndex].Values[currentValuesIndex] += number
// 	}
// 	return nil
// }

// type parser interface {
// 	nextChar(char rune)
// }

// func ParseMotions(fullPath string) (Path, error) {

// 	var currentParser parser

// 	// var motions Path
// 	// currentMotionIndex := -1
// 	// currentVectorIndex := -1
// 	// decimalIndex := 0
// 	// currentValuesIndex := 0
// 	// isParsingAdigit := false
// 	// isNegative := false

// 	for pos, char := range fullPath {
// 		switch {

// 		case unicode.IsLetter(char):
// 			motion, err := motion.InitMotion(char)

// 			if err != nil {
// 				return nil, err
// 			}
// 			motions = append(motions, motion)

// 			currentMotionIndex++
// 			currentValuesIndex = 0
// 			isParsingAdigit = false
// 			decimalIndex = 0

// 		case unicode.IsDigit(char):
// 			isParsingAdigit = true
// 			err := incrementCurrentPathPoint(char, &motions, currentMotionIndex, currentValuesIndex, pos, decimalIndex, isNegative)

// 			if err != nil {
// 				return nil, err
// 			}

// 			if decimalIndex > 0 {
// 				decimalIndex++
// 			}

// 		case char == '.': // char is period
// 			decimalIndex++

// 		case char == '-':
// 			isNegative = true

// 		case char == ' ' || char == ',': // char is space or ,
// 			decimalIndex = 0
// 			isNegative = false

// 			if isParsingAdigit {
// 				currentValuesIndex++
// 			}

// 		default:
// 			return nil, fmt.Errorf("invalid character %c at position %v", char, pos)
// 		}
// 	}
// 	return motions, nil
// }

// func splitMotions (path string) []string {
// 	var	unParsedMotions []string
// 	motionIndex := -1
// 	for _, char := range path {
// 		if (unicode.IsLetter(char)) {
// 			motionIndex++
// 			unParsedMotions = append(unParsedMotions, "")
// 		}
// 		unParsedMotions[motionIndex] += string(char)
// 	}

// 	return unParsedMotions
// }

// func ParseMotions(fullPath string) (Path, error) {
// 	splitedPath := splitMotions(fullPath)
// }

// could be a binary search if array was sorted

// func StringifyMotions(motions *Path) string {
// 	var output strings.Builder

// 	for i, motion := range *motions {
// 		output.WriteString(string(motion.Letter))
// 		for _, val := range motion.Values {
// 			output.WriteString(strconv.FormatFloat(val, 'f', -1, 64))
// 			output.WriteString(string(','))
// 		}

// 		if i < len(*motions)-1 {
// 			output.WriteString(string(' '))
// 		}

// 	}
// 	return output.String()
// }
