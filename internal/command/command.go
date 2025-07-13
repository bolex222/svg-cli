package command

import (
	"errors"

	"github.com/bolex222/svg-cli/internal/vector"
)

type CommandChar rune

const (
	MoveTo_golbal                    CommandChar = 'M'
	MoveTo_relative                  CommandChar = 'm'
	LineTo_global                    CommandChar = 'L'
	LineTo_relative                  CommandChar = 'l'
	HorizontalLineTo_global          CommandChar = 'H'
	HorizontalLineTo_relative        CommandChar = 'h'
	VerticalLineTo_global            CommandChar = 'V'
	VerticalLineTo_relative          CommandChar = 'v'
	CubicBezierTo_global             CommandChar = 'C'
	CubicBezierTo_relative           CommandChar = 'c'
	SmoothCubicBezierTo_global       CommandChar = 'S'
	SmoothCubicBezierTo_relative     CommandChar = 's'
	QuadraticBezierTo_global         CommandChar = 'Q'
	QuadraticBezierTo_relative       CommandChar = 'q'
	SmoothQuadraticBezierTo_global   CommandChar = 'T'
	SmoothQuadraticBezierTo_relative CommandChar = 't'
	ElipticalArcCurve_global         CommandChar = 'A'
	ElipticalArcCurve_relative       CommandChar = 'a'
	ClosePath_global                 CommandChar = 'Z'
	ClosePath_relative               CommandChar = 'z'
)

type CommandType int

const (
	NoValueCommand CommandType = iota
	SingleValueCommand
	DoubleValueCommand
	TripleValueCommand
	SpecialValueCommand
)

type Command struct {
	LargeArcFlag bool
	SweepFlag    bool
	Type         CommandType
	Letter       CommandChar
	Angle        float64
	Values       []vector.Vector2
}

func IsCharAValidCommand(char rune) bool {
	if char == rune(MoveTo_golbal) ||
		char == rune(MoveTo_relative) ||
		char == rune(LineTo_global) ||
		char == rune(LineTo_relative) ||
		char == rune(HorizontalLineTo_global) ||
		char == rune(HorizontalLineTo_relative) ||
		char == rune(VerticalLineTo_global) ||
		char == rune(VerticalLineTo_relative) ||
		char == rune(CubicBezierTo_global) ||
		char == rune(CubicBezierTo_relative) ||
		char == rune(SmoothCubicBezierTo_global) ||
		char == rune(SmoothCubicBezierTo_relative) ||
		char == rune(QuadraticBezierTo_global) ||
		char == rune(QuadraticBezierTo_relative) ||
		char == rune(SmoothQuadraticBezierTo_global) ||
		char == rune(SmoothQuadraticBezierTo_relative) ||
		char == rune(ElipticalArcCurve_global) ||
		char == rune(ElipticalArcCurve_relative) ||
		char == rune(ClosePath_global) ||
		char == rune(ClosePath_relative) {
		return true
	}
	return false
}

func isNoneValue(char rune) bool {
	commandChar := CommandChar(char)
	if commandChar == ClosePath_global || commandChar == ClosePath_relative {
		return true
	}
	return false
}

func isCharSingleValue(char rune) bool {
	commandChar := CommandChar(char)
	if commandChar == MoveTo_golbal ||
		commandChar == MoveTo_relative ||
		commandChar == LineTo_global ||
		commandChar == LineTo_relative ||
		commandChar == VerticalLineTo_global ||
		commandChar == VerticalLineTo_relative ||
		commandChar == HorizontalLineTo_global ||
		commandChar == HorizontalLineTo_relative ||
		commandChar == SmoothQuadraticBezierTo_global ||
		commandChar == SmoothQuadraticBezierTo_relative {
		return true
	}
	return false
}

func isCharDoubleValue(char rune) bool {
	commandChar := CommandChar(char)
	if commandChar == SmoothCubicBezierTo_global ||
		commandChar == SmoothCubicBezierTo_relative ||
		commandChar == QuadraticBezierTo_global ||
		commandChar == QuadraticBezierTo_relative {
		return true
	}
	return false
}

func isCharTripleValue(char rune) bool {
	commandChar := CommandChar(char)
	if commandChar == CubicBezierTo_global ||
		commandChar == CubicBezierTo_relative {
		return true
	}
	return false
}

func isCharSpecialValue(char rune) bool {
	commandChar := CommandChar(char)
	if commandChar == ElipticalArcCurve_global ||
		commandChar == ElipticalArcCurve_relative {
		return true
	}
	return false
}

func InitCommandFromChar(char rune) (*Command, error) {
	switch {
	case isNoneValue(char):
		return &Command{
			Letter: CommandChar(char),
			Type:   NoValueCommand,
		}, nil
	case isCharSingleValue(char):
		return &Command{
			Letter: CommandChar(char),
			Type:   NoValueCommand,
			Values: make([]vector.Vector2, 0, 1),
		}, nil
	case isCharDoubleValue(char):
		return &Command{
			Letter: CommandChar(char),
			Type:   NoValueCommand,
			Values: make([]vector.Vector2, 0, 2),
		}, nil
	case isCharTripleValue(char):
		return &Command{
			Letter: CommandChar(char),
			Type:   NoValueCommand,
			Values: make([]vector.Vector2, 0, 3),
		}, nil
	case isCharSpecialValue(char):
		return &Command{
			Letter: CommandChar(char),
			Type:   NoValueCommand,
			Values: make([]vector.Vector2, 0, 2),
		}, nil
	default:
		return nil, errors.New("")
	}
}
