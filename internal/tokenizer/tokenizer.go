package tokenizer

import (
	"bufio"
	"fmt"
	"unicode"
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
