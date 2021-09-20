package tv

// Конкретная команда
type SwitchCommand struct {
	Device     Device
	ChannelNum int
}

func (sc *SwitchCommand) Execute() {
	sc.Device.SwitchChannelTo(sc.ChannelNum)
}
