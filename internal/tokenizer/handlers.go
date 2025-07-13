package tokenizer

import (
	"errors"
	"fmt"
	"github.com/bolex222/svg-cli/internal/command"
	"strings"
	"unicode"
)

func (t *Tokenizer) handleDotChar(char rune) error {
	if !t.isCurrentlyNumberToken() || strings.Contains(t.currentToken.Value, ".") {
		t.finishCurrentToken()
		t.nextToken(TokenNumber, string(char))
	} else if err := t.appendToCurrentToken('.'); err != nil {
		return err
	}
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
