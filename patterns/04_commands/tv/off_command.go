package tv

type OffCommand struct {
	Device Device
}

func (oc *OffCommand) Execute() {
	oc.Device.Off()
}
