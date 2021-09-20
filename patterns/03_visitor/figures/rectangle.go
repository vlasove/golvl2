package figures

type Rectangle struct {
	A, B int
}

func NewRectangle(a, b int) *Rectangle {
	return &Rectangle{
		A: a,
		B: b,
	}
}

func (r *Rectangle) Accept(v Visitor) {
	v.VisitForRectangle(r)
}

func (r *Rectangle) GetType() string {
	return "Rectangle"
}
