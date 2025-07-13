package vector

type Vector2 struct {
	X float64
	Y float64
}

func New(x float64, y float64) Vector2 {
	return Vector2{X: x, Y: y}
}

func (v *Vector2) Add(other Vector2) {
	v.X += other.X
	v.Y += other.Y
}

func (v Vector2) Added(other Vector2) Vector2 {
	return Vector2{X: v.X + other.X, Y: v.Y + other.Y}
}
