package hospital

// Интерфейс обработчика
type Department interface {
	Execute(*Patient)
	SetNext(Department)
}
