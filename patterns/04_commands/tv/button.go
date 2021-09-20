package tv

// Конкретный отправитель (кнопка на панели телевизора)
type Button struct {
	Command Command
}

func (b *Button) Press() {
	b.Command.Execute()
}
