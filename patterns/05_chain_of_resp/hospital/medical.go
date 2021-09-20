package hospital

// Конкретный обработчик - кабинет медикаментов
type Medical struct {
	next Department
}

func NewMedical(next Department) *Medical {
	return &Medical{
		next: next,
	}
}

func (m *Medical) SetNext(next Department) {
	m.next = next
}

func (m *Medical) Execute(p *Patient) {
	m.accept(generalVisitor, p)
	m.next.Execute(p)
}

func (m *Medical) accept(v Visitor, p *Patient) {
	v.visitForMedical(p)
}
