package tokenizer

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
	"unicode"
	"github.com/bolex222/svg-cli/internal/command"
)

type TokenType int

const (
	TokenCommand TokenType = iota
	TokenNumber
	TokenEOF
	TokenError
)

type Token struct {
	Type  TokenType
	Value string
}

type Tokenizer struct {
	reader       *bufio.Reader
	currentToken *Token
	tokens       []Token
}

func New() *Tokenizer {
	return &Tokenizer{
		tokens:       make([]Token, 0),
		currentToken: nil,
	}
}

func (t *Tokenizer) nextToken(tokenType TokenType, initialString string) {
	t.currentToken = &Token{
		Type:  tokenType,
		Value: initialString,
	}
}

func (t *Tokenizer) finishCurrentToken() {
	if t.currentToken != nil {
		t.tokens = append(t.tokens, *t.currentToken)
		t.currentToken = nil
	}
}

func (t *Tokenizer) isCurrentlyNumberToken() bool {
	return t.currentToken != nil && t.currentToken.Type == TokenNumber
}

func (t *Tokenizer) appendToCurrentToken(char rune) error {
	if t.currentToken == nil {
		return fmt.Errorf("no token to append to")
	}

	t.currentToken.Value += string(char)
	return nil
}

func (t *Tokenizer) handleCharE(char rune) error {
	if !t.isCurrentlyNumberToken() {
		return errors.New("number token can not start by \"e\"")
	}
	if strings.Contains(t.currentToken.Value, "e") {
		return errors.New("token already contain a \"e\"")
	}
	if err := t.appendToCurrentToken(char); err != nil {
		return err
	}
	return nil
}

func (t *Tokenizer) handleMinusChar(char rune) error {
	p, err := t.reader.Peek(1)
	if err != nil {
		return errors.New("character \"-\" can not be a last char")
	}
	if t.currentToken != nil && t.currentToken.Type == TokenNumber && len(t.currentToken.Value) > 0 {
		lastLetter := t.currentToken.Value[len(t.currentToken.Value)-1]
		if lastLetter == 'e' {
			err := t.appendToCurrentToken(char)
			if err != nil {
				return err
			}
			return nil
		}
	}

	if !unicode.IsNumber(rune(p[0])) {
		return errors.New("character \"-\" can not be isolated")
	}
	t.finishCurrentToken()
	t.nextToken(TokenNumber, string(char))
	return nil
}

func (t *Tokenizer) handleDotChar(char rune) error {
	if !t.isCurrentlyNumberToken() || strings.Contains(t.currentToken.Value, ".") {
		t.finishCurrentToken()
		t.nextToken(TokenNumber, string(char))
	} else if err := t.appendToCurrentToken('.'); err != nil {
		return err
	}
	return nil
}

func (t *Tokenizer) handleNumberChar(char rune) error {
	if !t.isCurrentlyNumberToken() {
		t.finishCurrentToken()
		t.nextToken(TokenNumber, string(char))
	} else if err := t.appendToCurrentToken(char); err != nil {
		return err
	}
	return nil
}

func (t *Tokenizer) handleLetterChar(char rune) error {
	if !command.IsCharAValidCommand(char) {
		return fmt.Errorf("character %c is not a valid command", char)
	}
	t.finishCurrentToken()
	t.nextToken(TokenCommand, string(char))
	return nil
}

func (t *Tokenizer) Tokenize(reader *bufio.Reader) ([]Token, error) {
	t.reader = reader
	i := 0

	for {

		b, err := reader.ReadByte()
		if err != nil {
			t.finishCurrentToken()
			return t.tokens, nil
		}

		char := rune(b)

		switch {
		case char == 'e':
			err := t.handleCharE(char)
			if err != nil {
				return t.tokens, err
			}

		case char == '-':
			err := t.handleMinusChar(char)
			if err != nil {
				return t.tokens, err
			}

		case char == '.':
			err := t.handleDotChar(char)
			if err != nil {
				return t.tokens, err
			}

		case unicode.IsNumber(char):
			err := t.handleNumberChar(char)
			if err != nil {
				return t.tokens, err
			}

		case unicode.IsLetter(char):
			err := t.handleLetterChar(char)
			if err != nil {
				return t.tokens, err
			}

		case char == ' ' || char == ',':
			t.finishCurrentToken()

		default:
			return t.tokens, fmt.Errorf("unexpected character \"%c\" at pos %v", char, i+1)
		}
		i++
	}
}
