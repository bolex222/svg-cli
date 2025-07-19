package parser

import (
	"errors"

	"github.com/bolex222/svg-cli/internal/command"
	"github.com/bolex222/svg-cli/internal/lexer"
)

type CommandParser interface {
	PushToken(token lexer.Token, parser *Parser) error
	IsCommandOver() bool
	GetCommand() command.Command
}

type Parser struct {
	BufferCommand CommandParser
	Commands      []command.Command
}

func New() *Parser {
	return &Parser{
		BufferCommand: nil,
		Commands:      make([]command.Command, 0),
	}
}

func (p *Parser) getLastCommand() (command.Command, error) {
	if len(p.Commands) < 1 {
		return command.Command{}, errors.New("ther is no command")
	}

	return p.Commands[len(p.Commands)-1], nil
}

func selectParser(char command.CommandChar) (CommandParser, error) {
	switch char {
	case command.ElipticalArcCurve_relative, command.ElipticalArcCurve_global:
		elipticParser, err := NewElipticArcParser(char)
		if err != nil {
			return nil, err
		}
		return elipticParser, nil
	default:
		classicParser, err := NewClassicCommandParser(char)
		if err != nil {
			return nil, err
		}
		return classicParser, nil
	}
}

func (p *Parser) prepareCommandParser(token lexer.Token) error {
	if token.Type == lexer.TokenCommand {
		if len(token.Value) > 0 && command.IsCharAValidCommand(rune(token.Value[0])) {
			commandChar := command.CommandChar(token.Value[0])
			commandParser, err := selectParser(commandChar)
			if err != nil {
				return err
			}
			p.BufferCommand = commandParser
			return nil
		}
		return errors.New("invalid token")
	} else {
		previousCommand, err := p.getLastCommand()
		if err != nil {
			return errors.New("svg path must start with a Move to segment")
		}
		if command.IsCharAValidCommand(rune(previousCommand.Letter)) {
			var commandChar command.CommandChar
			switch previousCommand.Letter {
			case command.MoveTo_golbal:
				commandChar = command.LineTo_global
			case command.MoveTo_relative:
				commandChar = command.LineTo_relative
			default:
				commandChar = previousCommand.Letter
			}
			commandParser, err := selectParser(commandChar)
			if err != nil {
				return err
			}
			p.BufferCommand = commandParser
		}
		return nil
	}
}

func (p *Parser) commitCurrentCommand() {
	if p.BufferCommand != nil {
		p.Commands = append(p.Commands, p.BufferCommand.GetCommand())
		p.BufferCommand = nil
	}
}

func (p *Parser) ParseTokensToCommands(tokens []lexer.Token) ([]command.Command, error) {
	for _, token := range tokens {
		if p.BufferCommand == nil || token.Type == lexer.TokenCommand {
			err := p.prepareCommandParser(token)
			if err != nil {
				p.BufferCommand = nil
				return p.Commands, errors.New("parsing stopped due to invalid token")
			}
		}
		if token.Type == lexer.TokenNumber {
			err := p.BufferCommand.PushToken(token, p)
			if err != nil {
				p.BufferCommand = nil
				return p.Commands, errors.New("parsing stopped due to invalid token")
			}
		}
		if p.BufferCommand.IsCommandOver() {
			p.commitCurrentCommand()
		}
	}
	return p.Commands, nil
}
