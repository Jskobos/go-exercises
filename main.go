package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	durationPtr := flag.Int("duration", 30, "a duration")
	flag.Parse()
	timer := NewTimer(time.Duration(*durationPtr) * time.Second)
	game := NewGame()
	go timer.Run()
	go Play(game)
Mainloop:
	for {
		select {
		case msg := <-timer.TimerChannel:
			fmt.Println("Time's up!", msg)
			break Mainloop
		case msg1 := <-game.GameChannel:
			fmt.Println(msg1)
		}
	}
	EndGame(game)
}
