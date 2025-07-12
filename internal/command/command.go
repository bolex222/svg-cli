package command

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

// type Command struct {
// 	LargeArcFlag bool
// 	SweepFlag    bool
// 	Letter       rune
// 	Angle        float64
// 	Values       []vector.Vector2
// }
