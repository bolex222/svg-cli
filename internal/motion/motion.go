package motion

import (
	"fmt"

	"github.com/bolex222/svg-cli/internal/vector"
)

type Motion struct {
	LargeArcFlag bool
	SweepFlag    bool
	Letter       rune
	Angle        float64
	Values       []vector.Vector2
}

func InitMotion(letter rune) (Motion, error) {
	switch letter {
	case 'M':
		return Motion{
			Letter: letter,
			Values: make([]vector.Vector2, 0, 1),
		}, nil
	case 'm':
		return Motion{
			Letter: letter,
			Values: make([]vector.Vector2, 0, 1),
		}, nil
	case 'L':
		return Motion{
			Letter: letter,
			Values: make([]vector.Vector2, 0, 1),
		}, nil
	case 'l':

		return Motion{
			Letter: letter,
			Values: make([]vector.Vector2, 0, 1),
		}, nil
	case 'H':
		return Motion{
			Letter: letter,
			Values: make([]vector.Vector2, 0, 1),
		}, nil
	case 'h':
		return Motion{
			Letter: letter,
			Values: make([]vector.Vector2, 0, 1),
		}, nil
	case 'V':
		return Motion{
			Letter: letter,
			Values: make([]vector.Vector2, 0, 1),
		}, nil
	case 'v':
		return Motion{
			Letter: letter,
			Values: make([]vector.Vector2, 0, 1),
		}, nil
	case 'C':
		return Motion{
			Letter: letter,
			Values: make([]vector.Vector2, 0, 3),
		}, nil
	case 'c':
		return Motion{
			Letter: letter,
			Values: make([]vector.Vector2, 0, 3),
		}, nil
	case 'S':
		return Motion{
			Letter: letter,
			Values: make([]vector.Vector2, 0, 2),
		}, nil
	case 's':
		return Motion{
			Letter: letter,
			Values: make([]vector.Vector2, 0, 2),
		}, nil
	case 'Q':
		return Motion{
			Letter: letter,
			Values: make([]vector.Vector2, 0, 2),
		}, nil
	case 'q':
		return Motion{
			Letter: letter,
			Values: make([]vector.Vector2, 0, 2),
		}, nil
	case 'T':
		return Motion{
			Letter: letter,
			Values: make([]vector.Vector2, 0, 1),
		}, nil
	case 't':
		return Motion{
			Letter: letter,
			Values: make([]vector.Vector2, 0, 1),
		}, nil
	case 'A':
		return Motion{
			Letter: letter,
			Values: make([]vector.Vector2, 0, 2),
		}, nil
	case 'a':
		return Motion{
			Letter: letter,
			Values: make([]vector.Vector2, 0, 2),
		}, nil
	case 'Z':
		return Motion{
			Letter: letter,
		}, nil
	case 'z':
		return Motion{
			Letter: letter,
		}, nil
	default:
		return Motion{}, fmt.Errorf("Unexpected Letter %v \n", letter)
	}
}

