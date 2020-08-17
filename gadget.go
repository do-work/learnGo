package main

import (
	"github.com/learnGo/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func playlist(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	}
}

func main() {
	//mixtape := []string{"song1", "song2"}
	//var player Player = TapePlayer{}
	//playlist(player, mixtape)
	//player = TapeRecorder{}
	//playlist(player, mixtape)
	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})

}
