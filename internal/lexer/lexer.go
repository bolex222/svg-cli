package lexer

import (
	"bufio"
	"fmt"
	"unicode"
)

// newline char supported by W3C XML specification
// https://www.w3.org/TR/2008/REC-xml-20081126/#sec-line-ends
const CARRIAGE_RETURN = 0x0d
const LINE_FEED = 0x0a

type TokenType int

const (
	TokenCommand TokenType = iota
	TokenNumber
	TokenEOF   // NOT USED YET
	TokenError // NOT USED YET
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	reader       *bufio.Reader
	currentToken *Token
	tokens       []Token
}

func New() *Lexer {
	return &Lexer{
		tokens:       make([]Token, 0),
		currentToken: nil,
	}
}

func (t *Lexer) nextToken(tokenType TokenType, initialString string) {
	t.currentToken = &Token{
		Type:  tokenType,
		Value: initialString,
	}
}

func (t *Lexer) finishCurrentToken() {
	if t.currentToken != nil {
		t.tokens = append(t.tokens, *t.currentToken)
		t.currentToken = nil
	}
}

func (t *Lexer) isCurrentlyNumberToken() bool {
	return t.currentToken != nil && t.currentToken.Type == TokenNumber
}

func (t *Lexer) appendToCurrentToken(char rune) error {
	if t.currentToken == nil {
		return fmt.Errorf("no token to append to")
	}

	t.currentToken.Value += string(char)
	return nil
}

func (t *Lexer) Tokenize(reader *bufio.Reader) ([]Token, error) {
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

		case char == ' ' || char == ',' || char == CARRIAGE_RETURN || char == LINE_FEED:
			t.finishCurrentToken()

		default:
			return t.tokens, fmt.Errorf("unexpected character \"%c\" at pos %v", char, i+1)
		}
		i++
	}
}
