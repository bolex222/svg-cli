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
	r          *bufio.Reader
	tokens     []Token
	tokenIndex int
}

func New() Tokenizer {
	return Tokenizer{
		tokens:     make([]Token, 0),
		tokenIndex: -1,
	}
}

func (t *Tokenizer) nextToken(tokenType TokenType, initialString string) {
	t.tokenIndex++
	t.tokens = append(t.tokens, Token{
		Type:  tokenType,
		Value: initialString,
	})
}

func (t *Tokenizer) isCurrentlyNumberToken() bool {
	return t.tokenIndex >= len(t.tokens)-1 && t.tokens[t.tokenIndex].Type == TokenNumber
}

func (t *Tokenizer) Tokenize(r *bufio.Reader) ([]Token, error) {
	t.r = r
	i := 0

	for {
		b, err := r.ReadByte()
		if err != nil {
			return t.tokens, nil
		}
		char := rune(b)

		switch {
		case char == 'e':
			if !t.isCurrentlyNumberToken() {
				return t.tokens, fmt.Errorf("number token can not start by \"e\" at position %v", i)
			}

			if strings.Contains(t.tokens[t.tokenIndex].Value, "e") {
				return t.tokens, fmt.Errorf("token already contain a \"e\" at positoin %v", i)
			}

			t.tokens[t.tokenIndex].Value += string(char)

		case char == '-':
			p, err := r.Peek(1)
			if err != nil {
				return t.tokens, fmt.Errorf("character \"-\" can not be a last char")
			}

			if !unicode.IsNumber(rune(p[0])) {
				return t.tokens, fmt.Errorf("character \"-\" can not be isolated")
			}

			t.nextToken(TokenNumber, string(char))

		case char == '.':
			if !t.isCurrentlyNumberToken() || strings.Contains(t.tokens[t.tokenIndex].Value, ".") {
				t.nextToken(TokenNumber, string(char))
			} else {
				t.tokens[t.tokenIndex].Value += "."
			}

		case unicode.IsNumber(char):
			if !t.isCurrentlyNumberToken(){
				t.nextToken(TokenNumber, string(char))
			} else {
				t.tokens[t.tokenIndex].Value += string(char)
			}

		case unicode.IsLetter(char):
			// TODO: check if command is a valid letter
			t.nextToken(TokenCommand, string(char))

		case char == ' ' || char == ',':
			p, err := r.Peek(1)
			if err != nil {
				return t.tokens, nil
			}
			pChar := rune(p[0])

			if unicode.IsNumber(pChar) {
				t.nextToken(TokenNumber, "")
			}

		default:
			return t.tokens, fmt.Errorf("unexpected character \"%c\" at pos %v", char, i +1)
		}
		i++
	}
}
