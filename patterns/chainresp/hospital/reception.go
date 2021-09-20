package hospital

// Конкретный обработчик - применое отделение
type Reception struct {
	next Department
}

func NewReception(next Department) *Reception {
	return &Reception{
		next: next,
	}
}

func (r *Reception) SetNext(next Department) {
	r.next = next
}

func (r *Reception) Execute(p *Patient) {
	r.accept(generalVisitor, p)
	r.next.Execute(p)
}

func (r *Reception) accept(v Visitor, p *Patient) {
	v.visitForReception(p)
}
