package gadget

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
	var player Player = TapePlayer{}
	playlist(player, mixtape)
	player = TapeRecorder{}
	playlist(player, mixtape)
}
