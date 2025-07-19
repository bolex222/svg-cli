package parser

import (
	"fmt"
	"strconv"

	"github.com/bolex222/svg-cli/internal/command"
	"github.com/bolex222/svg-cli/internal/lexer"
	"github.com/bolex222/svg-cli/internal/vector"
)

type ClassicCommandParser struct {
	commandBuffer     command.Command
	maxTokenAmount    int
	currentTokenIndex int
}

func NewClassicCommandParser(char command.CommandChar) (CommandParser, error) {
	commandChar, err := command.InitCommandFromChar(char)
	if err != nil {
		return nil, err
	}

	return &ClassicCommandParser{
		commandBuffer:     *commandChar,
		maxTokenAmount:    int(commandChar.Type),
		currentTokenIndex: 0,
	}, nil
}

/*
* Fill current command with Token
 */
func (cp *ClassicCommandParser) PushToken(token lexer.Token, parser *Parser) error {
	if token.Type != lexer.TokenNumber {
		return fmt.Errorf("unexpected token type")
	}

	valuesIndex := cp.currentTokenIndex / 2
	valueElement := cp.currentTokenIndex % 2
	cp.currentTokenIndex++
	if valueElement == 0 {
		cp.commandBuffer.Values = append(cp.commandBuffer.Values, vector.New(0, 0))
	}

	if len(cp.commandBuffer.Values) < valuesIndex {
		return fmt.Errorf("unexpected error")
	}

	parsedValue, err := strconv.ParseFloat(token.Value, 64)
	if err != nil {
		return fmt.Errorf("token %v, could not be converted to valid floating number", token.Value)
	}

	/* handle case where the command is vertical or horizontal lineTo
	*  by using previous command missing value so the vector is complete
	*  for linear algebra transfomations
	 */
	if cp.commandBuffer.Type == command.HalfValueCommand {
		var isVertical = cp.commandBuffer.Letter == command.VerticalLineTo_global || cp.commandBuffer.Letter == command.VerticalLineTo_relative
		if isVertical {
			cp.commandBuffer.Values[valuesIndex].Y = parsedValue
			previousCommand, err := parser.getLastCommand()
			if err == nil && len(previousCommand.Values) > 0 {
				lastCommandValue := previousCommand.Values[len(previousCommand.Values)-1]
				cp.commandBuffer.Values[valuesIndex].X = lastCommandValue.X
			} else {
				cp.commandBuffer.Values[valuesIndex].X = 0
			}
		} else {
			cp.commandBuffer.Values[valuesIndex].X = parsedValue
			previousCommand, err := parser.getLastCommand()
			if err == nil && len(previousCommand.Values) > 0 {
				lastCommandValue := previousCommand.Values[len(previousCommand.Values)-1]
				cp.commandBuffer.Values[valuesIndex].Y = lastCommandValue.X
			} else {
				cp.commandBuffer.Values[valuesIndex].Y = 0
			}
		}
		// handle classic vector values
	} else {
		if valueElement == 1 {
			cp.commandBuffer.Values[valuesIndex].Y = parsedValue
		} else {
			cp.commandBuffer.Values[valuesIndex].X = parsedValue
		}
	}
	return nil
}

func (cp *ClassicCommandParser) IsCommandOver() bool {
	return cp.currentTokenIndex >= cp.maxTokenAmount
}

func (cp *ClassicCommandParser) GetCommand() command.Command {
	return cp.commandBuffer
}
