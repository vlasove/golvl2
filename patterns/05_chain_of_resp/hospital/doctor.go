package hospital

// Конкретный обработчик - осмотр доктора
type Doctor struct {
	next Department
}

func NewDoctor(next Department) *Doctor {
	return &Doctor{
		next: next,
	}
}

func (d *Doctor) SetNext(next Department) {
	d.next = next
}

func (d *Doctor) Execute(p *Patient) {
	d.accept(generalVisitor, p)
	d.next.Execute(p)
}

func (d *Doctor) accept(v Visitor, p *Patient) {
	v.visitForDoctor(p)
}
