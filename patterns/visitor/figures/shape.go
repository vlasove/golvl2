package figures

type Shape interface {
	GetType() string
	Accept(Visitor)
}
