package main

import "github.com/do-work/learnGo/gadget"

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

func main() {
	mixtape := []string{"song1", "song2"}
	var player Player = gadget.TapePlayer{}
	playlist(player, mixtape)
	player = gadget.TapeRecorder{}
	playlist(player, mixtape)
}
