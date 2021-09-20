package main

import (
	"log"

	"github.com/vlasove/materials/tasks_2/patterns/commands/cafe"
	"github.com/vlasove/materials/tasks_2/patterns/commands/tv"
)

// Очередь добавлена!

func main() {
	televisor := &tv.TV{}

	onCommand := &tv.OnCommand{
		Device: televisor,
	}

	offCommand := &tv.OffCommand{
		Device: televisor,
	}
	switchCommand := &tv.SwitchCommand{
		Device:     televisor,
		ChannelNum: 5,
	}

	onButton := &tv.Button{
		Command: onCommand,
	}
	onButton.Press()

	switchButton := &tv.Button{
		Command: switchCommand,
	}
	switchButton.Press()

	offButton := &tv.Button{
		Command: offCommand,
	}
	offButton.Press()

	switchCommand.ChannelNum = 10
	switchButton.Press()
	log.Println("now trying press remote buttons")

	onRemote := &tv.Remote{
		Command: onCommand,
	}

	onRemote.Press()

	offRemote := &tv.Remote{
		Command: offCommand,
	}
	offRemote.Press()

	// очередь
	coffix := cafe.New(25, 15)

	queue := []cafe.Command{
		coffix.MakeCapuccino(3),
		coffix.MakeLatte(5),
		coffix.MakeCapuccino(2),
		coffix.MakeClean(),
		coffix.MakeCapuccino(1),
	}

	for _, q := range queue {
		q.Execute()
	}
}
