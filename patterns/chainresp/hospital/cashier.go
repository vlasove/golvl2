package hospital

type Cashier struct {
	next Department
}

func NewCashier(next Department) *Cashier {
	return &Cashier{
		next: next,
	}
}

func (c *Cashier) SetNext(next Department) {
	c.next = next
}

func (c *Cashier) Execute(p *Patient) {
	c.accept(generalVisitor, p)
}

func (c *Cashier) accept(v Visitor, p *Patient) {
	v.visitForCashier(p)
}
