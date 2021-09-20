package tv

// Конкретный отправитель
type Remote struct {
	Command Command
}

func (r *Remote) Press() {
	r.Command.Execute()
}
