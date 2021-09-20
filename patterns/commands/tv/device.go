package tv

// Интерфейс получателя
type Device interface {
	On()
	Off()
	SwitchChannelTo(int)
}
