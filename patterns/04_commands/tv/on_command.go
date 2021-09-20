package tv

// Конкретная команда
type OnCommand struct {
	Device Device
}

func (oc *OnCommand) Execute() {
	oc.Device.On()
}
