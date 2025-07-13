package parser

import (
	"errors"

	"github.com/bolex222/svg-cli/internal/command"
	"github.com/bolex222/svg-cli/internal/tokenizer"
)

type Parser struct {
	BufferCommand *command.Command
	Commands      []command.Command
}

func New() *Parser {
	return &Parser{
		BufferCommand: nil,
		Commands:      make([]command.Command, 0),
	}
}

func (p *Parser) nextCommand(token tokenizer.Token) error {
	if token.Type == tokenizer.TokenCommand && len(token.Value) > 0 {
		char := rune(token.Value[0])
		command, err := command.InitCommandFromChar(char)
		if err != nil {
			return err
		}
		p.BufferCommand = command
		return nil
	} else {
		return errors.New("invalid token command")
	}
}

func (p *Parser) commitCommand() {
	if p.BufferCommand != nil {
		p.Commands = append(p.Commands, *p.BufferCommand)
	}
}

//func (p *Parser) handleNumberToken(token *tokenizer.Token) {
//	//TODO: define what happens in command token
//}

func (p *Parser) ParseTokensToCommands(tokens []tokenizer.Token) ([]command.Command, error) {
	for _, token := range tokens {
		switch token.Type {
		case tokenizer.TokenCommand:
			p.commitCommand()
			err := p.nextCommand(token)
			if err != nil {
				return p.Commands, err
			}
		default:
			continue
			// return p.Commands, nil
		}
	}
	p.commitCommand()
	return p.Commands, nil
}
