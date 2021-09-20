package figures

type Square struct {
	Side int
}

func NewSquare(s int) *Square {
	return &Square{
		Side: s,
	}
}

func (s *Square) Accept(v Visitor) {
	v.VisitForSquare(s)
}

func (s *Square) GetType() string {
	return "Square"
}
