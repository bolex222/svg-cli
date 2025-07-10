package vector

type Vector2 struct {
	X float64
	Y float64
}

func (v *Vector2) Add(other Vector2) {
	v.X += other.X
	v.Y += other.Y
}
