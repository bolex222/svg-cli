package tokenizer

import (
	"bufio"
	"fmt"
	"strings"
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
	reader            *bufio.Reader
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
			if !t.isCurrentlyNumberToken() {
				return t.tokens, fmt.Errorf("number token can not start by \"e\" at position %v", i)
			}

			if strings.Contains(t.currentToken.Value, "e") {
				return t.tokens, fmt.Errorf("token already contain a \"e\" at positoin %v", i)
			}

			if err := t.appendToCurrentToken(char); err != nil {
				return t.tokens, err
			}

		case char == '-':
			p, err := reader.Peek(1)
			if err != nil {
				return t.tokens, fmt.Errorf("character \"-\" can not be a last char")
			}

			if !unicode.IsNumber(rune(p[0])) {
				return t.tokens, fmt.Errorf("character \"-\" can not be isolated")
			}

			t.finishCurrentToken()
			t.nextToken(TokenNumber, string(char))

		case char == '.':
			if !t.isCurrentlyNumberToken() || strings.Contains(t.currentToken.Value, ".") {
				t.finishCurrentToken()
				t.nextToken(TokenNumber, string(char))
			} else if err := t.appendToCurrentToken('.'); err != nil {
				return t.tokens, err
			}

		case unicode.IsNumber(char):
			if !t.isCurrentlyNumberToken() {
				t.finishCurrentToken()
				t.nextToken(TokenNumber, string(char))
			} else if err := t.appendToCurrentToken(char); err != nil {
				return t.tokens, err
			}

		case unicode.IsLetter(char):
			// TODO: check if command is a valid letter
			t.finishCurrentToken()
			t.nextToken(TokenCommand, string(char))

		case char == ' ' || char == ',':
			t.finishCurrentToken()

		default:
			return t.tokens, fmt.Errorf("unexpected character \"%c\" at pos %v", char, i+1)
		}
		i++
	}
}
