package parser

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/bolex222/svg-cli/internal/command"
	"github.com/bolex222/svg-cli/internal/lexer"
	"github.com/bolex222/svg-cli/internal/vector"
)

type ElipticArcParser struct {
	commandBuffer      command.Command
	maxTokenAmount     int
	currentTokenIndex  int
	currentVectorIndex int
}

func NewElipticArcParser(char command.CommandChar) (CommandParser, error) {
	commandChar, err := command.InitCommandFromChar(char)
	if err != nil {
		return nil, err
	}

	return &ElipticArcParser{
		commandBuffer:      *commandChar,
		maxTokenAmount:     int(commandChar.Type),
		currentTokenIndex:  0,
		currentVectorIndex: 0,
	}, nil
}

const (
	angleTokenIndex int = 2
	largeArcFlag    int = 3
	sweepFlag       int = 4
)

/*
* Fill current command with Token
 */
func (cp *ElipticArcParser) PushToken(token lexer.Token, parser *Parser) error {
	if token.Type != lexer.TokenNumber {
		return fmt.Errorf("unexpected token type")
	}

	currentTokenIndex := cp.currentTokenIndex
	cp.currentTokenIndex++
	switch currentTokenIndex {
	case angleTokenIndex:
		value, err := strconv.ParseFloat(token.Value, 64)
		if err != nil {
			return errors.New("invalid token")
		}
		cp.commandBuffer.Angle = value
	case largeArcFlag, sweepFlag:
		value, err := strconv.ParseBool(token.Value)
		if err != nil {
			return errors.New("invalid token")
		}
		if currentTokenIndex == largeArcFlag {
			cp.commandBuffer.LargeArcFlag = value
		} else {
			cp.commandBuffer.SweepFlag = value
		}
	default:
		valuesIndex := cp.currentVectorIndex / 2
		valueElement := cp.currentVectorIndex % 2
		cp.currentVectorIndex++
		parsedValue, err := strconv.ParseFloat(token.Value, 64)
		if err != nil {
			return fmt.Errorf("token %v, could not be converted to valid floating number", token.Value)
		}
		if valueElement == 0 {
			cp.commandBuffer.Values = append(cp.commandBuffer.Values, vector.New(0, 0))
		}
		if len(cp.commandBuffer.Values) < valuesIndex {
			return fmt.Errorf("unexpected error")
		}
		if valueElement == 1 {
			cp.commandBuffer.Values[valuesIndex].Y = parsedValue
		} else {
			cp.commandBuffer.Values[valuesIndex].X = parsedValue
		}
	}
	return nil
}

func (cp *ElipticArcParser) IsCommandOver() bool {
	return cp.currentTokenIndex >= cp.maxTokenAmount
}

func (cp *ElipticArcParser) GetCommand() command.Command {
	return cp.commandBuffer
}
