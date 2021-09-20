package tv

import "log"

// Конкретный получатель
type TV struct {
	IsRunning bool
}

func (t *TV) On() {
	t.IsRunning = true
	log.Println("Turning On TV")
}

func (t *TV) Off() {
	t.IsRunning = false
	log.Println("Turning Off TV")
}

func (t *TV) SwitchChannelTo(channelNum int) {
	if t.IsRunning {
		log.Println("Now channel is switiching to:", channelNum)
		return
	}
	log.Println("Can not switch channel because TV is turned off. Wanted:", channelNum)
}
