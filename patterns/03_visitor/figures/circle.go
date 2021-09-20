package figures

type Circle struct {
	Radius float64
}

func NewCircle(r float64) *Circle {
	return &Circle{
		Radius: r,
	}
}

func (c *Circle) Accept(v Visitor) {
	v.VisitForCircle(c)
}

func (c *Circle) GetType() string {
	return "Circle"
}
